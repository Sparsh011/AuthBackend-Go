package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func OtpServiceClientId() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("OTP_SERVICE_CLIENT_ID")
}

func OtpServiceClientSecret() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("OTP_SERVICE_CLIENT_SECRET")
}
