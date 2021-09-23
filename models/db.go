package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	fmt.Println("Attempting to connect to 'firi' database...")
	err := godotenv.Load("conf.env")
	if err != nil {
		LogError(errors.New("error accessing config file"))
	}

	username, password, dbName, dbHost, err := getDatabaseCredentials()
	if err != nil {
		LogError(err)
	}

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn, err := gorm.Open("postgres", dbURL)
	if err != nil {
		LogError(err)
	}

	db = conn
	fmt.Println("Core database connection successful")
	// autoMigrateTables()
}

func getDatabaseCredentials() (string, string, string, string, error) {
	_ = godotenv.Load("conf.env")

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	return username, password, dbName, dbHost, nil

}

//GetDB sends the db objects
func GetDB() *gorm.DB {
	return db
}
