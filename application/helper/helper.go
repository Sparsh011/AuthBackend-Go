package helper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetErrorStatusCode(err error) int {
	var statusCode int
	if strings.Contains(err.Error(), "invalid parameters") {
		statusCode = http.StatusBadRequest // 400
	} else if strings.Contains(err.Error(), "unauthorized") {
		statusCode = http.StatusUnauthorized // 401
	} else if strings.Contains(err.Error(), "too many requests") {
		statusCode = http.StatusTooManyRequests // 429
	} else {
		statusCode = http.StatusInternalServerError // 500
	}

	return statusCode
}

func CreateJWTToken(secretKey []byte, identifier string, exp time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  identifier,
			"exp": exp.Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
