package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "Welcome!\n")
}

// /hello/sparsh -> gives output hello, sparsh
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Create a map to hold the raw JSON data
	var rawRequest map[string]interface{}

	// Decode the JSON request body into the map
	if err := json.NewDecoder(request.Body).Decode(&rawRequest); err != nil {
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

	// Process the data (e.g., send OTP)
	fmt.Fprintf(writer, "Send OTP to number: %s\n", phoneNumber)
	fmt.Fprintf(writer, "OTP Length: %d\n", int(otpLength))
	fmt.Fprintf(writer, "Timeout: %d seconds\n", int(timeout))

	// Handle extra fields if needed
	for key, value := range rawRequest {
		if key != PhoneNumberKey && key != OtpLengthKey && key != TimeoutKey {
			fmt.Fprintf(writer, "Extra field: %s = %v\n", key, value)
		}
	}

	response := map[string]interface{}{
		PhoneNumberKey: phoneNumber,
		OtpLengthKey:   int(otpLength),
		TimeoutKey:     int(timeout),
	}

	// Set response headers
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Encode the response map to JSON and write it to the response
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)
	// router.POST("/otp-login", SendOtp)
	// fmt.Println("Starting server at port 8080...")
	// log.Fatal(http.ListenAndServe(":8080", router))

	queryParams := url.Values{} // in `Values`, the type of values is slice of string and key is string. Whole `Values` is a map of string, slice
	queryParams.Add("country", "in")
	queryParams.Add("country", "us")
	queryParams.Add("category", "sports")
	queryParams.Add("apiKey", "6c2421d5b3174ab5b768cc14dbccc3ba")
	queryParams.Add("pageSize", "2")
	// ?country=in&category=sports&pageSize=2
	resp, err := GetRequestHandler(
		"https://newsapi.org",
		"/v2/top-headlines",
		queryParams,
		nil,
		nil,
	)

	if err != nil {
		fmt.Println("Error while making apiCall: ", err)
		return
	}

	fmt.Println("Response of news api call: ", resp)

}

type OtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	OtpLength   int    `json:"otpLength"`
	Timeout     int    `json:"timeout"`
}

const (
	PhoneNumberKey = "phoneNumber"
	OtpLengthKey   = "otpLength"
	TimeoutKey     = "timeout"
)
