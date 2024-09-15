package authhandler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
)

func RefreshToken(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Parse the request body to get the refresh token
	var req authpkg.AccessTokenRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newAccessToken, accessCreationErr := helper.CreateAccessTokenFromRefreshToken(req.RefreshToken)

	if accessCreationErr != nil || len(strings.TrimSpace(newAccessToken)) == 0 {
		http.Error(writer, "Unable to authorize user, please login again.", http.StatusUnauthorized)
		return
	}

	response := authpkg.AccessTokenResponse{AccessToken: newAccessToken}
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
	}
}
