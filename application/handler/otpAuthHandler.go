package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/data"
	"github.com/sparsh011/AuthBackend-Go/application/helper"
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
	timeout, timeoutExtractionError := rawRequest[TimeoutKey].(float64)

	if !phoneExtractionError || !otpLengthExtractionError || !timeoutExtractionError {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	headers := getOTPApiHeaders()

	otpRequestData := data.SendOtpRequest{
		PhoneNumber: phoneNumber,
		OtpLength:   int(otpLength),
		Channel:     "SMS",
		Expiry:      int(timeout),
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

	// Encode the response map to JSON. This returns a json back to the frontend/client
	err := json.NewEncoder(writer).Encode(sendOtpResponse)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
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

	if !phoneExtractionError || !otpExtractionError || !orderIdExtractionError {
		http.Error(writer, "Missing required fields", http.StatusBadRequest)
		return
	}

	headers := getOTPApiHeaders()

	verifyOtpRequest := data.VerifyOtpRequest{
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

	// Encode the response map to JSON. This returns a json back to the frontend/client
	err := json.NewEncoder(writer).Encode(verifyOtpResponse)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
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

	resendOtpRequest := data.ResendOtpRequest{
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

	// Encode the response map to JSON and write it to the response
	if jsonParsingError := json.NewEncoder(writer).Encode(resendOtpResponse); jsonParsingError != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
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
	TimeoutKey     = "timeout"
	OrderIdKey     = "orderId"
	OtpKey         = "otp"
)

const (
	OTPClientBaseURL = "https://auth.otpless.app"
	OTPSendRoute     = "/auth/otp/v1/send"
	OTPResendRoute   = "/auth/otp/v1/resend"
	OTPVerifyRoute   = "/auth/otp/v1/verify"
)
