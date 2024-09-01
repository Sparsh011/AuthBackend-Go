package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/routes"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

func main() {
	// Initialize DB before starting server
	service.InitializeDB()

	// Starting server
	router := httprouter.New()
	routes.ConfigureRoutesAndStartServer(router)
}
