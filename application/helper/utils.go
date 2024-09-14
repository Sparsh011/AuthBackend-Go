package helper

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetErrorStatusCode(err error) int {
	var statusCode = 400 // Set bad request as a fallback in case parsing fails
	if strings.Contains(err.Error(), "invalid parameters") {
		statusCode = http.StatusBadRequest // 400
	} else if strings.Contains(err.Error(), "unauthorized") {
		statusCode = http.StatusUnauthorized // 401
	} else if strings.Contains(err.Error(), "too many requests") {
		statusCode = http.StatusTooManyRequests // 429
	} else {
		statusCode = http.StatusInternalServerError // 500
	}

	return statusCode
}

func CreateUserVerificationResponse(
	isVerified bool,
	message string,
	access string,
	refresh string,
	userProfile map[string]interface{},
) map[string]interface{} {
	verificationResponse := make(map[string]interface{})

	verificationResponse["isVerified"] = isVerified
	verificationResponse["message"] = message
	verificationResponse["access"] = access
	verificationResponse["refresh"] = refresh
	verificationResponse["userProfile"] = userProfile

	return verificationResponse
}

func GetRandomName() string {
	superheroes := []string{
		"Superman", "Batman", "Wonder Woman", "Flash", "Green Lantern",
		"Aquaman", "Cyborg", "Spider-Man", "Iron Man", "Thor",
		"Captain America", "Black Widow", "Hulk", "Doctor Strange",
		"Black Panther", "Scarlet Witch", "Ant-Man", "Wolverine",
		"Deadpool", "Silver Surfer",
	}

	rand.Seed(time.Now().UnixNano())

	randomSuperhero := superheroes[rand.Intn(len(superheroes))]

	randomNumber := rand.Intn(901) + 100

	return fmt.Sprintf("%s-%d", randomSuperhero, randomNumber)
}
