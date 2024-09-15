package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/initializers"
	"github.com/sparsh011/AuthBackend-Go/application/routes"
)

func main() {
	// Load env file
	initializers.LoadEnvFile()

	// Initialize DB before starting servere
	initializers.InitializeDB()

	// Starting server
	router := httprouter.New()
	routes.ConfigureRoutesAndStartServer(router)
}
