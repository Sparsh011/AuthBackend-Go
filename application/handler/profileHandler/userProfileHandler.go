package profilehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/db"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
)

func UserProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authHeader := request.Header.Get("Authorization")

	authHeader = strings.TrimSpace(authHeader)

	if authHeader == "" {
		http.Error(writer, "Authorization header is missing.", http.StatusUnauthorized)
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		http.Error(writer, "Invalid Authorization header format.", http.StatusUnauthorized)
		return
	}

	isTokenValid, validationError := helper.ValidateJWT(strings.TrimSpace(authHeaderParts[1]))

	if !isTokenValid || validationError != nil {
		fmt.Println("Validation error: ", validationError.Error())
		http.Error(writer, "Invalid token.", http.StatusUnauthorized)
		return
	}

	userId, userIdExtractionErr := helper.ExtractUserID(authHeaderParts[1])

	if userIdExtractionErr != nil {
		http.Error(writer, "Unable to authorize user, please login again.", http.StatusUnauthorized)
		return
	}

	user, userFetchErr := db.FindUserByID(userId)

	if user == nil || userFetchErr != nil {
		http.Error(writer, "Could not find user details.", http.StatusInternalServerError)
		return
	}

	if json := json.NewEncoder(writer).Encode(user); json != nil {
		http.Error(writer, "Something went wrong!", http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
}
