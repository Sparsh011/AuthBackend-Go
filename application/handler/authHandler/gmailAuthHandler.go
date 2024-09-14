package authhandler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/db"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	"github.com/sparsh011/AuthBackend-Go/application/initializers"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
	"google.golang.org/api/idtoken"
)

type IdTokenRequest struct {
	Token string `json:"token"`
}

func ValidateGoogleIDTokenHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var tokenRequest IdTokenRequest
	if err := json.NewDecoder(request.Body).Decode(&tokenRequest); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the Google ID token
	tokenInfo, err := idtoken.Validate(context.Background(), tokenRequest.Token, initializers.GetGoogleWebClientID())
	if err != nil {
		http.Error(writer, "Invalid token", http.StatusUnauthorized)
		return
	}

	var userProfile map[string]interface{} = nil

	// If email is verified, only then create user profile
	if tokenInfo.Claims["email_verified"] == true {
		verificationTime := time.Now()
		userRandomName := helper.GetRandomName()
		userId := uuid.New()

		userProfile = map[string]interface{}{
			"verificationTime": verificationTime,
			"expenseBudget":    -1,
			"name":             userRandomName,
			"emailId":          tokenInfo.Claims["email"],
			"profileUri":       tokenInfo.Claims["picture"],
			"userId":           userId,
		}

		user := authpkg.User{
			VerificationTime: verificationTime,
			ExpenseBudget:    -1,
			Name:             userRandomName,
			EmailId:          sql.NullString{String: tokenInfo.Claims["email"].(string), Valid: true},
			ProfileUri:       tokenInfo.Claims["picture"].(string),
			Id:               userId,
		}

		isInserted, err := db.InsertUser(&user)

		if err != nil || !isInserted {
			http.Error(writer, "Unable to save user details.", http.StatusInternalServerError)
			return
		}
	}

	access, accessCreationError := helper.CreateJWTToken(
		[]byte(initializers.GetJWTSigningKey()),
		tokenInfo.Claims["email"].(string),
		time.Now().Add(time.Hour*48),
	)
	if accessCreationError != nil {
		http.Error(writer, "Unable to sign in.", http.StatusInternalServerError)
		return
	}

	refresh, refreshCreationError := helper.CreateJWTToken(
		[]byte(initializers.GetJWTSigningKey()),
		tokenInfo.Claims["email"].(string),
		time.Now().Add(10*365*24*time.Hour),
	)
	if refreshCreationError != nil {
		http.Error(writer, "Unable to sign in.", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"isVerified":  true,
		"message":     "Gmail verified successfully.",
		"access":      access,
		"refresh":     refresh,
		"userProfile": userProfile,
	}

	writer.Header().Set("Content-Type", "application/json")
	if jsonParsingError := json.NewEncoder(writer).Encode(response); jsonParsingError != nil {
		http.Error(writer, "Something went wrong!", http.StatusInternalServerError)
	}
}
