package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sparsh011/AuthBackend-Go/application/routes"
)

func main() {
	router := httprouter.New()
	routes.ConfigureRoutesAndStartServer(router)
}
