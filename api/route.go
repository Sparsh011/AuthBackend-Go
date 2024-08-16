package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Configures all the routes and starts the server
func Routes(router *httprouter.Router) {
	router.POST(SendOtpRoute, SendOtp)
	router.POST(ResendOtpRoute, ResendOtp)
	router.POST(VerifyOtpRoute, VerifyOtp)
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8000", router))
}

const (
	SendOtpRoute   = "/login/send-otp"
	ResendOtpRoute = "/login/resent-otp"
	VerifyOtpRoute = "/login/verify-otp"
)
