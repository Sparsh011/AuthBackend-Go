package service

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

func GetPort() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}

	return os.Getenv("PORT")
}

func GetDBPort() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_PORT")
}

func GetDBUsername() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_USERNAME")
}

func GetDBHost() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_HOST")
}

func GetDBName() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_NAME")
}

func GetDBPassword() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_PASSWORD")
}
