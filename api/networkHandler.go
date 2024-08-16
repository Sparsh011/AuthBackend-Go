package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
networkUrl: Base URL of the API
apiRoute: API route, e.g., "/users/{id}"
params: Query parameters, e.g., ?name=john&age=30
pathVars: Path variables, e.g., {id: "123"}
headers: Headers, e.g., {"Authorization": "Bearer token"}
*/
func GetRequestHandler(
	networkUrl string,
	apiRoute string,
	queryParams url.Values,
	pathVars map[string]string,
	headers map[string]string,
) (map[string]interface{}, error) {
	// Replace path variables in the route.
	// 1st: string in which replacement will be made
	// 2nd: string that will be used to replace a value
	// 3rd: string that will be replaced with the 2nd param
	// 4th: -1 indicated all occurrences should be replaced. ideally in our case, only 1 occurrence of a path variable should be there
	for key, value := range pathVars {
		apiRoute = strings.Replace(apiRoute, "{"+key+"}", value, -1)
	}

	// Construct the full URL with query parameters
	fullURL := fmt.Sprintf("%s%s?%s", networkUrl, apiRoute, queryParams.Encode())

	// Create the GET request
	req, getRequestCreationError := http.NewRequest("GET", fullURL, nil)
	if getRequestCreationError != nil {
		return nil, fmt.Errorf("error creating request: %v", getRequestCreationError)
	}

	// Add headers to the request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Make the GET request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	responseData, readingResponseError := io.ReadAll(resp.Body)
	if readingResponseError != nil {
		return nil, fmt.Errorf("error reading response: %v", readingResponseError)
	}

	// Unmarshal the response body into a map
	var jsonResponse map[string]interface{}
	if unmarshallingError := json.Unmarshal(responseData, &jsonResponse); unmarshallingError != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", unmarshallingError)
	}

	// Return the JSON response
	return jsonResponse, nil
}
