package app

import (
	"log"
	"main/domain"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	eh := EventHandlers{service.NewEventService(domain.NewEventRepositoryStub())}

	//define routes
	router.HandleFunc("/events", eh.getAllEvents).Methods(http.MethodGet)
	// router.HandleFunc("/events", eh.postAnEvent).Methods(http.MethodPost)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
