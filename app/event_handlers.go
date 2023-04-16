package app

import (
	"encoding/json"
	"main/dto"
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
func (eh *EventHandlers) createEvent(w http.ResponseWriter, r *http.Request) {
	var request dto.NewEventRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		event, error := eh.EventService.CreateEvent(request)
		if error != nil {
			writeResponse(w, 403, error)
		} else {
			writeResponse(w, http.StatusCreated, event)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
