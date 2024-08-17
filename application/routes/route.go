package routes

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/handler"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

// Configures all the routes and starts the server
func ConfigureRoutesAndStartServer(router *httprouter.Router) {
	const (
		SendOtpRoute   = "/login/send-otp"
		ResendOtpRoute = "/login/resent-otp"
		VerifyOtpRoute = "/login/verify-otp"
		HomeRoute      = "/"
	)

	router.POST(SendOtpRoute, handler.SendOtp)
	router.POST(ResendOtpRoute, handler.ResendOtp)
	router.POST(VerifyOtpRoute, handler.VerifyOtp)
	log.Fatal(http.ListenAndServe(service.GetPort(), router))
}
