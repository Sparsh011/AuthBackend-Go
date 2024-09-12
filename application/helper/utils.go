package helper

import (
	"net/http"
	"strings"
)

func GetErrorStatusCode(err error) int {
	var statusCode = 400 // Set bad request as a fallback in case parsing fails
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
