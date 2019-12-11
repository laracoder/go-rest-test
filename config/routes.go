package config

import (
	"fmt"
	"github.com/laracoder/go-rest-api/event"
	"net/http"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var ApiRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		HomeLink,
	},
	Route{
		"EventList",
		"GET",
		"/events",
		event.GetAllEvents,
	},
	Route{
		"Event",
		"GET",
		"/events/{id}",
		event.GetEvent,
	},
}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("event-rest-service")
}
