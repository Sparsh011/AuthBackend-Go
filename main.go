package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/routes"
	"github.com/sparsh011/AuthBackend-Go/application/service"
)

func main() {
	router := httprouter.New()
	// Initialize DB before starting server
	service.InitializeDB()
	routes.ConfigureRoutesAndStartServer(router)
}
