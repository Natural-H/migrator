package main

import (
	"dummyMigration/InitUtils"
	"dummyMigration/models"
	"flag"
	"fmt"
	"github.com/eiannone/keyboard"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
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

	var credentials, err = InitUtils.ReadCredentialsFromFile("./credentials.json")

	if err != nil {
		fmt.Println("Error reading credentials from file, initializing new credentials...")
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

	fmt.Println("Database migrated successfully.\nNow Mocking data...")

	models.MockDB(db)

	fmt.Println("Done!\nPress any key to continue...")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	_, _, _ = keyboard.GetKey()
	fmt.Println("Key pressed, continuing...")
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
