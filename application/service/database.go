package service

import (
	"errors"
	"fmt"

	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
	"github.com/sparsh011/AuthBackend-Go/application/models/expense"
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
	DB.AutoMigrate(&authpkg.User{}, &expense.Expense{})
}

func InsertUser(user *authpkg.User) (bool, error) {
	var existingUser authpkg.User
	var result *gorm.DB

	if !user.PhoneNumber.Valid || user.PhoneNumber.String != "" {
		result = DB.Where("phone_number = ?", user.PhoneNumber).First(&existingUser)
	} else if !user.EmailId.Valid || user.EmailId.String != "" {
		result = DB.Where("email_id = ?", user.EmailId).First(&existingUser)
	} else {
		return false, errors.New("either phone number or email ID must be provided")
	}

	if result.Error == nil {
		// User exists, update the LastLoginTime time
		existingUser.VerificationTime = user.VerificationTime
		updateResult := DB.Save(&existingUser)
		if updateResult.Error != nil {
			return false, updateResult.Error
		}
		return true, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		// User doesn't exist, create a new user
		createResult := DB.Create(&user)
		if createResult.Error != nil {
			return false, createResult.Error
		}
		return true, nil
	} else {
		// Some other error occurred
		return false, result.Error
	}
}
