package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sparsh011/AuthBackend-Go/application/initializers"
)

var secretKey = []byte(initializers.GetJWTSigningKey())

// CreateJWTToken generates a new JWT token using userId and an expiration time
func CreateJWTToken(userId string, exp time.Time) (string, error) {
	// Create a new token with claims
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = jwt.NewNumericDate(exp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAuthorizationHeader(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the expiration time in milliseconds
		if exp, ok := claims["exp"].(float64); ok {
			// Convert milliseconds to time.Time
			expTime := int64(exp) // Unix timestamp in seconds

			// Convert the Unix timestamp to time.Time
			expirationTime := time.Unix(expTime, 0)

			// Check if the token is expired
			if time.Now().After(expirationTime) {
				return false, fmt.Errorf("token is expired")
			}

			// Token is valid and not expired
			return true, nil
		}
		return false, fmt.Errorf("exp claim is missing or invalid")
	}

	return false, fmt.Errorf("invalid token")
}

func ExtractUserID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid and extract the userId
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	}

	return "", fmt.Errorf("invalid token")
}

func CreateAccessTokenFromRefreshToken(refreshToken string) (string, error) {
	userId, err := ExtractUserID(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %v", err)
	}

	expirationTime := time.Now().Add(time.Hour * 48)
	return CreateJWTToken(userId, expirationTime)
}

func ExtractTokenFromHeader(authHeader string) (string, error) {
	if len(strings.TrimSpace(authHeader)) == 0 {
		return "", fmt.Errorf("unable to authorize")
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("unable to authorize")
	}

	return authHeaderParts[1], nil
}
