package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/handler"
	authhandler "github.com/sparsh011/AuthBackend-Go/application/handler/authHandler"
	profilehandler "github.com/sparsh011/AuthBackend-Go/application/handler/profileHandler"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

func ConfigureRoutesAndStartServer(router *httprouter.Router) {
	const (
		SendOtpRoute      = "/login/send-otp"
		ResendOtpRoute    = "/login/resend-otp"
		VerifyOtpRoute    = "/login/verify-otp"
		RefreshTokenRoute = "/user/refresh"
		UserProfileRoute  = "/user/profile"
		VerifyTokenRoute  = "/login/otp/verify-token"
		HomeRoute         = "/"
	)

	router.POST(SendOtpRoute, authhandler.SendOtp)
	router.POST(ResendOtpRoute, authhandler.ResendOtp)
	router.POST(VerifyOtpRoute, authhandler.VerifyOtp)
	router.POST(VerifyTokenRoute, authhandler.ValidateOtpVerificationToken)
	router.GET(HomeRoute, handler.IndexHandler)
	router.POST(RefreshTokenRoute, authhandler.RefreshToken)
	router.GET(UserProfileRoute, profilehandler.UserProfile)

	fmt.Println("Starting server at port", service.GetPort())

	err := http.ListenAndServe(service.GetPort(), router)
	if err != nil {
		println("Error starting server:", err.Error())
		log.Fatal(err)
	}
}
