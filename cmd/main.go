package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/api"
)

// /hello/sparsh -> gives output hello, sparsh
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	api.Routes(router)

	// resp, err := api.GetRequestHandler(
	// 	"https://newsapi.org",
	// 	"/v2/top-headlines",
	// 	queryParams,
	// 	nil,
	// 	nil,
	// )

	// if err != nil {
	// 	fmt.Println("Error while making apiCall: ", err)
	// 	return
	// }

	// fmt.Println("Response of news api call: ", resp)
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
