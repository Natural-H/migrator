package main

import (
	"bufio"
	"dummyMigration/InitUtils"
	"dummyMigration/models"
	"flag"
	"fmt"
	"github.com/eiannone/keyboard"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
)

func createAndWriteCredentialsFile() InitUtils.DatabaseCredentials {
	credentials := InitUtils.CreateCredentials()

	fmt.Printf("This credentials will be written to credentials.json (WARN: Password is in plain text)")
	err := InitUtils.WriteCredentialsToFile("./credentials.json", credentials)
	if err != nil {
		log.Fatalf("Error writing credentials to file: %v", err)
	}

	return credentials
}

func main() {
	reset := flag.Bool("reset", false, "Reset the database credentials")
	flag.Parse()

	fmt.Println("Trying to read credentials from file...")
	var credentials, err = InitUtils.ReadCredentialsFromFile("./credentials.json")

	if err != nil {
		fmt.Println("Couldn't read credentials from file, initializing new credentials...")
		credentials = createAndWriteCredentialsFile()
	} else {
		if *reset {
			fmt.Println("Resetting credentials...")
			credentials = createAndWriteCredentialsFile()
		} else {
			fmt.Println("Credentials loaded from file.")
		}
	}

	_, _, _ = keyboard.GetKey()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", credentials.User, credentials.Password, credentials.Host, credentials.Port, credentials.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		SkipDefaultTransaction: credentials.SkipDefaultTransaction,
		PrepareStmt:            !credentials.DontGeneratePreparedStatements,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = DropAllTables(db)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = db.AutoMigrate(
		models.Libros{},
		models.Usuarios{},
		models.Prestamos{},
		models.Autores{},
		models.Generos{},
		models.Editoriales{},
		models.LibrosAutores{},
		models.LibrosGeneros{},
	)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			log.Fatalf("failed to close keyboard: %v", err)
		}
	}()

	fmt.Println("Database migrated successfully.\nMock data? (Y/n)")
	key := "y"
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		key = input
	}

	//_, key, _ := keyboard.GetKey()

	if key == "y" || key == "Y" {
		MockData(db, reader)
	} else {
		fmt.Println("Skipping mock data generation...")
	}

	fmt.Println("Done!\nPress any key to continue...")

	_, _, _ = keyboard.GetKey()
}

func MockData(db *gorm.DB, reader *bufio.Reader) {
	var registers = 100
	var prestamos = 1500

	fmt.Print("Enter how many Users (Usuarios) will be created (default: 100)> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		_, err := fmt.Sscanf(input, "%d", &registers)
		if err != nil {
			log.Fatalf("Error user count: %s", err)
		}
	}

	fmt.Print("Enter how many Lends (Préstamos) will be created (default: 1500)> ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		_, err := fmt.Sscanf(input, "%d", &registers)
		if err != nil {
			log.Fatalf("Error lend count: %s", err)
		}
	}

	fmt.Println("Generating mock data...")
	models.MockDB(db, registers, prestamos)
}

func DropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return err
	}

	for _, table := range tables {
		if err = db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}
