package main

import (
	"github.com/gorilla/mux"
	"github.com/laracoder/go-rest-api/services"
	"log"
	"net/http"
	"os"
)

var (
	serverPort = os.Getenv("PORT")
)

func main() {
	if serverPort == "" {
		log.Fatal("$PORT must be set")
		os.Exit(1)
	}
	log.Println("Application server started at port " + serverPort)

	var router = getRouter()
	services.RegisterApiRoutes(router)

	log.Fatal(http.ListenAndServe(":"+serverPort, router))
}

func getRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
