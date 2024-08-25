package routes

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/handler"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

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
	router.GET(HomeRoute, handler.IndexHandler)

	err := http.ListenAndServe(service.GetPort(), router)
	if err != nil {
		println("Error starting server:", err.Error())
		log.Fatal(err)
	}
}
