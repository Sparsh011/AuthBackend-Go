package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	godotenv.Unmarshal(".env")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}
}

func OtpServiceClientId() string {
	return os.Getenv("OTP_SERVICE_CLIENT_ID")
}

func OtpServiceClientSecret() string {
	return os.Getenv("OTP_SERVICE_CLIENT_SECRET")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetDBPort() string {
	return os.Getenv("DB_PORT")
}

func GetDBUsername() string {
	return os.Getenv("DB_USERNAME")
}

func GetDBHost() string {
	return os.Getenv("DB_HOST")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetJWTSigningKey() string {
	return os.Getenv("JWT_SIGNING_KEY")
}

func GetGoogleWebClientID() string {
	return os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
}
