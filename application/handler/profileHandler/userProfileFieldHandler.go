package profilehandler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/db"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
)

type UpdateProfileFieldRequest struct {
	NewValue string `json:"newValue"`
	Field    string `json:"field"`
}

const (
	ProfileUri    = "profile_uri"
	Name          = "name"
	ExpenseBudget = "expense_budget"
)

func UpdateUserProfileField(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	if len(strings.TrimSpace(userId)) == 0 || userIdExtractionErr != nil {
		http.Error(writer, "Unable to authorize", http.StatusUnauthorized)
		return
	}

	var updateRequest UpdateProfileFieldRequest
	decodingErr := json.NewDecoder(request.Body).Decode(&updateRequest)

	if decodingErr != nil {
		http.Error(writer, "Invalid request body.", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	switch updateRequest.Field {
	case ProfileUri:
		{
			savingErr := db.UpdateProfileUri(userId, updateRequest.NewValue)
			handleErrAndResponse(writer, savingErr, updateRequest.Field)
			return
		}
	case Name:
		{
			savingErr := db.UpdateUserName(userId, updateRequest.NewValue)
			handleErrAndResponse(writer, savingErr, updateRequest.Field)
			return
		}
	case ExpenseBudget:
		{
			num64, numberParsingErr := strconv.Atoi(updateRequest.NewValue)
			if numberParsingErr != nil {
				http.Error(writer, "Invalid expense budget.", http.StatusBadRequest)
				return
			}

			num32 := int32(num64)
			savingErr := db.UpdateUserExpenseBudget(userId, num32)
			handleErrAndResponse(writer, savingErr, updateRequest.Field)
			return
		}
	default:
		{
			http.Error(writer, "Unexpected field.", http.StatusBadRequest)
			return
		}

	}
}

func handleErrAndResponse(writer http.ResponseWriter, savingErr error, field string) {
	if savingErr != nil {
		http.Error(writer, "Unable to update fields, please try again.", http.StatusInternalServerError)
		return
	}

	responseJson := map[string]interface{}{
		"isSuccess":    true,
		"field":        field,
		"errorMessage": nil,
	}

	parsingErr := json.NewEncoder(writer).Encode(responseJson)

	if parsingErr != nil {
		http.Error(writer, "Unable to parse request", http.StatusInternalServerError)
	}
}
