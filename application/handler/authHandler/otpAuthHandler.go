package authhandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
	authpkg "github.com/sparsh011/AuthBackend-Go/application/models/authPkg"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

func SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Create a map to hold the raw JSON data
	var rawRequest map[string]interface{}

	// Decode the JSON request body into the map
	if bodyParsingError := json.NewDecoder(request.Body).Decode(&rawRequest); bodyParsingError != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract known fields
	phoneNumber, phoneExtractionError := rawRequest[PhoneNumberKey].(string)
	otpLength, otpLengthExtractionError := rawRequest[OtpLengthKey].(float64) // JSON numbers are float64 by default
	expiry, expiryExtractionError := rawRequest[ExpiryKey].(float64)

	if !phoneExtractionError || !otpLengthExtractionError || !expiryExtractionError {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	if int(expiry) < 60 {
		http.Error(writer, "Expiry cannot be less than 60 seconds.", http.StatusBadRequest)
	}

	headers := getOTPApiHeaders()

	otpRequestData := authpkg.SendOtpRequest{
		PhoneNumber: phoneNumber,
		OtpLength:   int(otpLength),
		Channel:     "SMS",
		Expiry:      int(expiry),
	}

	body, otpRequestParsingError := json.Marshal(otpRequestData)
	if otpRequestParsingError != nil {
		http.Error(writer, "Error marshaling json", http.StatusInternalServerError)
		return
	}

	sendOtpResponse, sendOtpError := helper.PostRequestHandler(
		OTPClientBaseURL,
		OTPSendRoute,
		nil,
		headers,
		body,
	)

	if sendOtpError != nil {
		http.Error(writer, sendOtpError.Error(), helper.GetErrorStatusCode(sendOtpError))
		return
	}

	writer.WriteHeader(http.StatusOK)
	// Encode the response map to JSON. This returns a json back to the frontend/client
	err := json.NewEncoder(writer).Encode(sendOtpResponse)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
}

func VerifyOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Create a map to hold the raw JSON data
	var rawRequest map[string]interface{}

	// Decode the JSON request body into the map
	if bodyParsingError := json.NewDecoder(request.Body).Decode(&rawRequest); bodyParsingError != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract known fields
	phoneNumber, phoneExtractionError := rawRequest[PhoneNumberKey].(string)
	otp, otpExtractionError := rawRequest[OtpKey].(string)
	orderId, orderIdExtractionError := rawRequest[OrderIdKey].(string)

	println("Phone number: ", phoneNumber, "\nOTP: ", otp, "\norderId: ", orderId)

	if !phoneExtractionError || !otpExtractionError || !orderIdExtractionError {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	headers := getOTPApiHeaders()

	verifyOtpRequest := authpkg.VerifyOtpRequest{
		PhoneNumber: phoneNumber,
		Otp:         otp,
		OrderId:     orderId,
	}

	body, bodyParsingError := json.Marshal(verifyOtpRequest)

	if bodyParsingError != nil {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	verifyOtpResponse, verifyOtpError := helper.PostRequestHandler(
		OTPClientBaseURL,
		OTPVerifyRoute,
		nil,
		headers,
		body,
	)

	if verifyOtpError != nil {
		http.Error(writer, verifyOtpError.Error(), helper.GetErrorStatusCode(verifyOtpError))
		return
	}
	if verifyOtpResponse["isOTPVerified"] != true {
		// Check if the "message" is present and not nil
		var message string
		if verifyOtpResponse["message"] != nil {
			message = verifyOtpResponse["message"].(string)
		} else if verifyOtpResponse["reason"] != nil {
			message = verifyOtpResponse["reason"].(string)
		} else {
			message = "Unknown error"
		}

		// Send the error response
		http.Error(writer, "OTP verification failed because "+message, http.StatusBadRequest)
		return
	}

	access, accessCreationError := helper.CreateJWTToken(
		[]byte(service.GetJWTSigningKey()),
		phoneNumber,
		time.Now().Add(time.Hour*48),
	)

	if accessCreationError != nil {
		http.Error(writer, "Failed to create access token.", http.StatusInternalServerError)
		return
	}

	refresh, refreshCreationError := helper.CreateJWTToken(
		[]byte(service.GetJWTSigningKey()),
		phoneNumber,
		time.Now().Add(10*365*24*time.Hour),
	)

	if refreshCreationError != nil {
		http.Error(writer, "Failed to create refresh token.", http.StatusInternalServerError)
		return
	}

	service.InsertUser(
		&authpkg.User{
			CreatedAt:     time.Now(),
			ExpenseBudget: 0,
			PhoneNumber:   phoneNumber,
		},
	)

	verifyOtpResponse["access"] = access
	verifyOtpResponse["refresh"] = refresh

	writer.WriteHeader(http.StatusOK)
	// Encode the response map to JSON. This returns a json back to the frontend/client
	err := json.NewEncoder(writer).Encode(verifyOtpResponse)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
}

func ResendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Create a map to hold the raw JSON data
	var rawRequest map[string]interface{}

	// Decode the JSON request body into the map
	if bodyParsingError := json.NewDecoder(request.Body).Decode(&rawRequest); bodyParsingError != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract known fields
	orderId, orderIdExtractionError := rawRequest[OrderIdKey].(string)

	if !orderIdExtractionError {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	headers := getOTPApiHeaders()

	resendOtpRequest := authpkg.ResendOtpRequest{
		OrderId: orderId,
	}

	body, bodyParsingError := json.Marshal(resendOtpRequest)

	if bodyParsingError != nil {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	resendOtpResponse, resendOtpError := helper.PostRequestHandler(
		OTPClientBaseURL,
		OTPResendRoute,
		nil,
		headers,
		body,
	)

	if resendOtpError != nil {
		http.Error(writer, resendOtpError.Error(), helper.GetErrorStatusCode(resendOtpError))
		return
	}

	writer.WriteHeader(http.StatusOK)
	// Encode the response map to JSON and write it to the response
	if jsonParsingError := json.NewEncoder(writer).Encode(resendOtpResponse); jsonParsingError != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
}

func getOTPApiHeaders() map[string]string {
	return map[string]string{
		"clientId":     service.OtpServiceClientId(),
		"clientSecret": service.OtpServiceClientSecret(),
		"Content-Type": "application/json",
	}
}

const (
	PhoneNumberKey = "phoneNumber"
	OtpLengthKey   = "otpLength"
	ExpiryKey      = "expiry"
	OrderIdKey     = "orderId"
	OtpKey         = "otp"
)

const (
	OTPClientBaseURL = "https://auth.otpless.app"
	OTPSendRoute     = "/auth/otp/v1/send"
	OTPResendRoute   = "/auth/otp/v1/resend"
	OTPVerifyRoute   = "/auth/otp/v1/verify"
)
