package db

import (
	"database/sql"
	"log"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// database creation with Migration of data
func Connect() *sql.DB {
	_ = godotenv.Load()
	file, err := os.Create(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	db, err := sql.Open(os.Getenv("DB_DRIVER"), "./"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	MigrationUp()

	log.Println("Database connection started")

	return db
}

// utilizing golang-migrate to load the data in sqllite db from the files
func MigrationUp() {
	_ = godotenv.Load()
	var driver database.Driver

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_NAME"))

	if err != nil {

		log.Fatalln(err)

	}

	if driver, err = sqlite.WithInstance(db, &sqlite.Config{}); err != nil {

		log.Fatalln(err)
	}
	migration, err := migrate.NewWithDatabaseInstance(
		"file://./db/dbfiles",
		os.Getenv("DB_DRIVER"), driver)
	if err != nil {
		log.Fatalln(err)
	}
	err = migration.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalln(err)
	}
}
