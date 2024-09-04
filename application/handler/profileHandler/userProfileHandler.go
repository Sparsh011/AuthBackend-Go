package profilehandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
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

	isTokenValid, validationError := helper.ValidateJWT(authHeaderParts[1])

	if !isTokenValid || validationError != nil {
		fmt.Println("Validation error: ", validationError.Error())
		http.Error(writer, "Invalid token.", http.StatusUnauthorized)
		return
	}

	// Get profile from postgres and send it back to client
}
