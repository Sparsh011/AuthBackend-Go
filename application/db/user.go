package db

import (
	"errors"

	"github.com/sparsh011/AuthBackend-Go/application/initializers"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
	"gorm.io/gorm"
)

func InsertUser(user *authpkg.User) (bool, error) {
	var existingUser authpkg.User
	var result *gorm.DB

	if !user.PhoneNumber.Valid || user.PhoneNumber.String != "" {
		result = initializers.DB.Where("phone_number = ?", user.PhoneNumber).First(&existingUser)
	} else if !user.EmailId.Valid || user.EmailId.String != "" {
		result = initializers.DB.Where("email_id = ?", user.EmailId).First(&existingUser)
	} else {
		return false, errors.New("either phone number or email ID must be provided")
	}

	if result.Error == nil {
		// User exists, update the LastLoginTime time
		existingUser.VerificationTime = user.VerificationTime
		updateResult := initializers.DB.Save(&existingUser)
		if updateResult.Error != nil {
			return false, updateResult.Error
		}
		return true, nil
	} else if result.Error == gorm.ErrRecordNotFound {
		// User doesn't exist, create a new user
		createResult := initializers.DB.Create(&user)
		if createResult.Error != nil {
			return false, createResult.Error
		}
		return true, nil
	} else {
		// Some other error occurred
		return false, result.Error
	}
}
