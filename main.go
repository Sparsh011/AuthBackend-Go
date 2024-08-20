package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/storage/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/models"
	"github.com/sparsh011/AuthBackend-Go/application/routes"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	router := httprouter.New()
	routes.ConfigureRoutesAndStartServer(router)

	db, dbError := postgres.New(postgres.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	}
	repository := Repository{
		DB: db,
	}
}

func (r *Repository) SaveUserDetails(context *fiber.Ctx) error {
	user := models.User{
		PhoneNumber:  "9643519024",
		CountryCode:  "91",
		AccessToken:  "access",
		RefreshToken: "refresh",
		LastLoginAt:  time.Now().UnixNano() / 1e6,
	}

	userParsingError := context.BodyParser(&user)

	if userParsingError != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"},
		)
		return userParsingError
	}

	userCreationError := r.DB.Create(&user).Error

	if userCreationError != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "could not create user"},
		)
		return userCreationError
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "User added successfully"},
	)

	return nil
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	userModels := &[]models.Users{}

	fetchUsersError := r.DB.Find(userModels).Error

	if fetchUsersError != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get users"},
		)
		return fetchUsersError
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Users successfully fetched",
		"data":    userModels,
	})

	return nil
}
