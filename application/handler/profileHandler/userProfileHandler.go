package profilehandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func UserProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// authHeader := request.Header.Get("Authorization")

	// authHeader = strings.TrimSpace(authHeader)

	// if authHeader == "" {
	// 	http.Error(writer, "Authorization header is missing.", http.StatusUnauthorized)
	// 	return
	// }

	// authHeaderParts := strings.Split(authHeader, " ")

	// if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
	// 	http.Error(writer, "Invalid Authorization header format.", http.StatusUnauthorized)
	// 	return
	// }

	// isTokenValid, validationError := helper.ValidateJWT(authHeaderParts[1])

	// if !isTokenValid || validationError != nil {
	// 	fmt.Println("Validation error: ", validationError.Error())
	// 	http.Error(writer, "Invalid token.", http.StatusUnauthorized)
	// 	return
	// }

	// Get profile from postgres and send it back to client
	dummyData := make(map[string]interface{})
	dummyData["createdAt"] = time.Now()
	dummyData["expenseBudget"] = 50000
	dummyData["name"] = "Sparsh"
	dummyData["phoneNumber"] = "1234567890"
	dummyData["email"] = "abc@gmail.com"
	dummyData["userId"] = "8d88e73d-66d3-4275-bc46-3534c8a74009"

	encodingError := json.NewEncoder(writer).Encode(dummyData)

	if encodingError != nil {
		http.Error(writer, encodingError.Error(), http.StatusInternalServerError)
	}
}
