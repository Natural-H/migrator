package main

import (
	"dummyMigration/models"
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type DatabaseCredentials struct {
	Host                           string `json:"host"`
	Port                           int    `json:"port"`
	User                           string `json:"user"`
	Password                       string `json:"password"`
	DBName                         string `json:"dbName"`
	DontGeneratePreparedStatements bool
	SkipDefaultTransaction         bool
}

func main() {
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")
	user := flag.String("user", "root", "Database user")
	password := flag.String("password", "", "Database password")
	database := flag.String("database", "biblioteca", "Database name")
	flag.Parse()

	credentials := DatabaseCredentials{
		Host:                           *host,
		Port:                           *port,
		User:                           *user,
		Password:                       *password,
		DBName:                         *database,
		SkipDefaultTransaction:         false,
		DontGeneratePreparedStatements: true,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", credentials.User, credentials.Password, credentials.Host, credentials.Port, credentials.DBName)
	var err error
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

	models.MockDB(db)
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
