package authhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

func RefreshToken(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Parse the request body to get the refresh token
	var req authpkg.AccessTokenRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse and validate the refresh token
	claims := &helper.JWTClaims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			signingError := "Unexpected signing error: " + token.Header["alg"].(string)
			return nil, fmt.Errorf(signingError, http.StatusUnauthorized)
		}
		return []byte(service.GetJWTSigningKey()), nil
	})

	if err != nil || !token.Valid {
		http.Error(writer, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	// Create a new access token
	expirationTime := time.Now().Add(time.Hour * 48) // Access token valid for 48 hours
	newAccessToken, err := helper.CreateJWTToken(
		[]byte(service.GetJWTSigningKey()),
		string(claims.UserID),
		expirationTime,
	)
	if err != nil {
		http.Error(writer, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Respond with the new access token
	response := authpkg.AccessTokenResponse{AccessToken: newAccessToken}
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
