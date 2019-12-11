package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/laracoder/go-rest-api/event"
	"log"
	"net/http"
)

const serverPort = "8080"

func main() {
	log.Println("Application server started at port " + serverPort)
	router := mux.NewRouter().StrictSlash(true)
	registerApiRoutes(router)
	log.Fatal(http.ListenAndServe(":" + serverPort, router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func registerApiRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", homeLink)
	api.HandleFunc("/events", event.GetAllEvents)
	api.HandleFunc("/events/{id}", event.GetEvent)
}

