package app

import (
	"encoding/json"
	"main/service"
	"net/http"
)

type EventHandlers struct {
	EventService service.EventService
}

// Get all Events from Event Service and return them as JSON
func (eh *EventHandlers) getAllEvents(w http.ResponseWriter, r *http.Request) {
	events, _ := eh.EventService.GetAllEvents()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// Create a new Event and return it as JSON
// func (eh *EventHandlers) postAnEvent(w http.ResponseWriter, r *http.Request) {
// 	event, _ := eh.EventService.CreateEvent()

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(event)
// }
