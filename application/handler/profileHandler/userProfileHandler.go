package profilehandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/db"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
)

func GetUserProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authHeader := request.Header.Get("Authorization")

	token, tokenExtractionErr := helper.ExtractTokenFromHeader(authHeader)

	if len(strings.TrimSpace(token)) == 0 || tokenExtractionErr != nil {
		http.Error(writer, "Unable to authorize", http.StatusUnauthorized)
		return
	}

	isTokenValid, validationError := helper.ValidateAuthorizationHeader(strings.TrimSpace(token))

	if !isTokenValid || validationError != nil {
		http.Error(writer, "Unable to authorize.", http.StatusUnauthorized)
		return
	}

	userId, userIdExtractionErr := helper.ExtractUserID(token)

	if userIdExtractionErr != nil {
		http.Error(writer, "Unable to authorize user.", http.StatusUnauthorized)
		return
	}

	user, userFetchErr := db.FindUserByID(userId)

	if user == nil || userFetchErr != nil {
		http.Error(writer, "Could not find user details.", http.StatusInternalServerError)
		return
	}

	userDto := authpkg.UserDto{
		VerificationTime: user.VerificationTime,
		ExpenseBudget:    user.ExpenseBudget,
		Name:             user.Name,
		PhoneNumber:      helper.HandleNullString(user.PhoneNumber),
		EmailId:          helper.HandleNullString(user.EmailId),
		ProfileUri:       user.ProfileUri,
	}

	if json := json.NewEncoder(writer).Encode(userDto); json != nil {
		http.Error(writer, "Something went wrong!", http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
}
