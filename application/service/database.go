package service

import (
	"fmt"
	"time"

	"github.com/sparsh011/AuthBackend-Go/application/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	println("Starting db donnection")

	var err error
	println("Starting db donnection")
	dbPort := GetDBPort()
	dbPassword := GetDBPassword()
	dbUser := GetDBUsername()
	dbName := GetDBName()
	dbHost := GetDBHost()

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=require", dbUser, dbPassword, dbName, dbHost, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to")
		return nil, err
	}

	return DB, nil
}

func SyncDatabase() {
	fmt.Println(
		"Suyncing DB")
	DB.AutoMigrate(&models.DBUser{})

	dummyUser := models.DBUser{
		Name:               "Sparsh chadha",
		PhoneNumber:        "55555555",
		CountryCode:        "+1",
		RefreshToken:       "dummy_refresh_token",
		LastLoginAt:        time.Now().Unix(),
		RefreshTokenExpiry: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}

	// Insert the dummy user into the database
	result := DB.Create(&dummyUser)

	print("Result: ", result)

}
