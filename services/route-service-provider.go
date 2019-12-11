package services

import (
	"github.com/gorilla/mux"
	"github.com/laracoder/go-rest-api/config"
	"log"
	"net/http"
)

const apiPrefix = "/api/v1"

func RegisterApiRoutes(router *mux.Router) {

	api := router.PathPrefix(apiPrefix).Subrouter()

	for _, route := range config.ApiRoutes {
		var handler http.Handler
		log.Println("registered route '" + route.Name + "'")
		handler = route.HandlerFunc

		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
