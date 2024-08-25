package service

import (
	"fmt"

	"github.com/sparsh011/AuthBackend-Go/application/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	ConnectToDatabase()
	SyncDatabase()
}

func ConnectToDatabase() (*gorm.DB, error) {
	var err error

	dbPort := GetDBPort()
	dbPassword := GetDBPassword()
	dbUser := GetDBUsername()
	dbName := GetDBName()
	dbHost := GetDBHost()

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=require", dbUser, dbPassword, dbName, dbHost, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		errorStr := "Failed to connect to DB, " + err.Error()
		panic(errorStr)
	}

	return DB, nil
}

// Perform AutoMigrate here to sync the schema with server.
// NOTE: AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
// It will change existing column's type if its size, precision, nullable changed.
// It WON'T delete unused columns to protect your data
func SyncDatabase() {
	DB.AutoMigrate(&models.ETUser{})
}
