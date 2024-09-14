package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/handler"
	authhandler "github.com/sparsh011/AuthBackend-Go/application/handler/authHandler"
	profilehandler "github.com/sparsh011/AuthBackend-Go/application/handler/profileHandler"
	"github.com/sparsh011/AuthBackend-Go/application/initializers"
)

func ConfigureRoutesAndStartServer(router *httprouter.Router) {
	const (
		HomeRoute              = "/"
		RefreshTokenRoute      = "/user/refresh"
		UserProfileRoute       = "/user/profile"
		VerifyTokenRoute       = "/login/otp/verify-token"
		VerifyGoogleOAuthToken = "/login/gmail/verify-token"
	)

	// Index route
	router.GET(HomeRoute, handler.IndexHandler)

	// User profile routes
	router.GET(UserProfileRoute, profilehandler.UserProfile)
	router.POST(RefreshTokenRoute, authhandler.RefreshToken)

	// Auth routes
	router.POST(VerifyTokenRoute, authhandler.ValidateOtpVerificationTokenHandler)
	router.POST(VerifyGoogleOAuthToken, authhandler.ValidateGoogleIDTokenHandler)

	fmt.Println("Starting server at port", initializers.GetPort())

	err := http.ListenAndServe(initializers.GetPort(), router)
	if err != nil {
		println("Error starting server:", err.Error())
		log.Fatal(err)
	}
}
