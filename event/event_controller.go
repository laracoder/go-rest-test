package event

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/laracoder/go-rest-api/event/types"
	"net/http"
)

const contentType = "application/json"

var events = types.AllEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "Introduction to Golang2",
		Description: "Testing",
	},
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			w.Header().Set("Content-Type", contentType)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)
	json.NewEncoder(w).Encode(events)
}