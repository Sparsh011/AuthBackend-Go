package authhandler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/db"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	"github.com/sparsh011/AuthBackend-Go/application/initializers"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
)

func ValidateOtpVerificationTokenHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var rawRequest map[string]interface{}

	if bodyParsingError := json.NewDecoder(request.Body).Decode(&rawRequest); bodyParsingError != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, tokenParsingError := rawRequest["token"].(string)

	if !tokenParsingError || len(strings.TrimSpace(token)) == 0 {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("client_id", initializers.OtpServiceClientId())
	formData.Set("client_secret", initializers.OtpServiceClientSecret())

	verifyTokenResponse, tokenVerificationError := helper.PostRequestHandler(
		VerifyTokenApiBaseURL,
		VerifyTokenApiRoute,
		nil,
		headers,
		formData,
		"application/x-www-form-urlencoded",
	)

	if tokenVerificationError != nil {
		http.Error(writer, tokenVerificationError.Error(), helper.GetErrorStatusCode(tokenVerificationError))
		return
	}

	phoneNumber, ok := verifyTokenResponse["phone_number"].(string)
	if !ok || len(strings.TrimSpace(phoneNumber)) == 0 {
		http.Error(writer, string("Invalid credentials."), http.StatusBadRequest)
		return
	}

	access, accessCreationError := helper.CreateJWTToken(
		[]byte(initializers.GetJWTSigningKey()),
		phoneNumber,
		time.Now().Add(time.Hour*48),
	)

	if accessCreationError != nil {
		http.Error(writer, "Failed to create access token.", http.StatusInternalServerError)
		return
	}

	refresh, refreshCreationError := helper.CreateJWTToken(
		[]byte(initializers.GetJWTSigningKey()),
		phoneNumber,
		time.Now().Add(10*365*24*time.Hour),
	)

	if refreshCreationError != nil {
		http.Error(writer, "Failed to create refresh token.", http.StatusInternalServerError)
		return
	}

	verificationTime := time.Now()
	userRandomName := helper.GetRandomName()
	userId := uuid.New()

	user := authpkg.User{
		VerificationTime: verificationTime,
		ExpenseBudget:    -1,
		Name:             userRandomName,
		PhoneNumber:      sql.NullString{String: phoneNumber, Valid: true},
		EmailId:          sql.NullString{},
		ProfileUri:       "",
		Id:               userId,
	}

	isInserted, err := db.InsertUser(&user)

	if err != nil || !isInserted {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userProfile := map[string]interface{}{
		"verifiedAt":    verificationTime,
		"expenseBudget": -1,
		"name":          userRandomName,
		"phoneNumber":   phoneNumber,
		"emailId":       "",
		"profileUri":    "",
		"userId":        userId,
	}

	jsonResponse := map[string]interface{}{
		"isVerified":  true,
		"message":     "Otp verification successful!",
		"access":      access,
		"refresh":     refresh,
		"userProfile": userProfile,
	}

	if jsonParsingError := json.NewEncoder(writer).Encode(jsonResponse); jsonParsingError != nil {
		http.Error(writer, jsonParsingError.Error(), http.StatusInternalServerError)
	}
}

const (
	VerifyTokenApiBaseURL = "https://auth.otpless.app"
	VerifyTokenApiRoute   = "/auth/userInfo"
)
