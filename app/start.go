package app

import (
	"database/sql"
	"fmt"
	"log"
	"main/domain"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	s := sqlClient()

	// wiring
	er := domain.NewEventRepositoryDb(&s)
	eh := EventHandlers{service.NewEventService(er)}

	// define routes
	router.HandleFunc("/events", eh.getAllEvents).Methods(http.MethodGet)
	router.HandleFunc("/events", eh.createEvent).Methods(http.MethodPost)

	// starting server
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func sqlClient() sql.DB {
	client, err := sql.Open("sqlite3", "event_log_dump")
	if err != nil {
		panic(err)
	}

	return *client
}
