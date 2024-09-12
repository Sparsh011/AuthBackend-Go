package authhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/idtoken"
)

type IdTokenRequest struct {
	IdToken string `json:"idToken"`
}

func ValidateGmailIDTokenHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var tokenRequest IdTokenRequest

	if err := json.NewDecoder(request.Body).Decode(&tokenRequest); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tokenInfo, err := idtoken.Validate(context.Background(), tokenRequest.IdToken, "YOUR_SERVER_CLIENT_ID")
	if err != nil {
		http.Error(writer, "Invalid token", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"email":   tokenInfo.Claims["email"],
		"sub":     tokenInfo.Subject,
	}

	writer.Header().Set("Content-Type", "application/json")
	jsonParsingError := json.NewEncoder(writer).Encode(response)

	if jsonParsingError != nil {
		http.Error(writer, "Something Went Wrong!", http.StatusInternalServerError)
	}
}
