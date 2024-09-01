package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

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

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.GetJWTSigningKey()), nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the expiration time
		exp, ok := claims["exp"].(float64)
		if !ok {
			return false, fmt.Errorf("invalid exp claim")
		}

		// Convert exp to time.Time
		expTime := time.Unix(int64(exp), 0)

		// Check if the token is expired
		return time.Now().After(expTime), nil
	}

	return false, fmt.Errorf("invalid token")
}

type JWTClaims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}
