package db

import (
	"errors"
	"fmt"
	"time"

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

func FindUserIfExists(email string, phone string) (*authpkg.User, error) {
	var user authpkg.User

	// Check if the user exists and fetch the user record
	err := initializers.DB.Where("email_id = ? OR phone_number = ?", email, phone).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User does not exist
			return nil, nil
		}
		// Some other error occurred
		return nil, err
	}

	return &user, nil
}

func UpdateVerificationTime(existingUser *authpkg.User, updatedVerificationTime time.Time) error {
	var user authpkg.User

	// Find the user by ID
	err := initializers.DB.First(&user, "id = ?", existingUser.Id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return err
	}

	fmt.Println("got user: ", user)

	// Update the user details with the new values (only specific fields)
	user.VerificationTime = updatedVerificationTime

	// Save the updated user details
	err = initializers.DB.Save(&user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func FindUserByID(userId string) (*authpkg.User, error) {
	var user authpkg.User

	err := initializers.DB.Where("id = ? ", userId).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func UpdateProfileUri(userId string, profileUri string) error {
	var user authpkg.User

	userFetchErr := initializers.DB.First(&user, "id = ?", userId).Error
	if userFetchErr != nil {
		if userFetchErr == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return userFetchErr
	}

	user.ProfileUri = profileUri

	savingErr := initializers.DB.Save(&user).Error
	if savingErr != nil {
		return fmt.Errorf("failed to update user: %v", savingErr)
	}

	return nil
}

func UpdateUserName(userId string, name string) error {
	var user *authpkg.User

	userFetchErr := initializers.DB.First(&user, "id = ?", userId).Error

	if userFetchErr != nil {
		if userFetchErr == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return userFetchErr
	}

	user.Name = name

	savingErr := initializers.DB.Save(&user).Error

	if savingErr != nil {
		return fmt.Errorf("failed to update user: %v", savingErr)
	}

	return nil
}

func UpdateUserExpenseBudget(userId string, budget int32) error {
	var user *authpkg.User

	userFetchErr := initializers.DB.First(&user, "id = ? ", userId).Error

	if userFetchErr != nil {
		if userFetchErr == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found")
		}
		return userFetchErr
	}

	user.ExpenseBudget = budget

	userSaveErr := initializers.DB.Save(user).Error

	if userSaveErr != nil {
		return fmt.Errorf("failed to update user: %v", userSaveErr)
	}

	return nil
}
