package domain

import (
	"database/sql"
	"log"
	"main/errs"

	_ "github.com/mattn/go-sqlite3"
)

type EventRepositoryDb struct {
	client *sql.DB
}

// Return all Events from the database
func (er EventRepositoryDb) GetAllEvents() ([]Event, *errs.AppError) {
	findAllSql := "select * from events"

	rows, err := er.client.Query(findAllSql)
	if err != nil {
		log.Fatal("Error while querying events table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	events := make([]Event, 0)
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.EventType, &e.CreatedAt, &e.Metadata)
		if err != nil {
			log.Fatal("Error while scanning events" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected error from database")
		}
		events = append(events, e)
	}

	return events, nil
}

// Create a new Event to the database
func (er EventRepositoryDb) CreateEvent(e Event) (*Event, *errs.AppError) {
	sqlInsert := "INSERT INTO events (event_type, created_at, metadata) values (?, ?, ?)"

	result, err := er.client.Exec(sqlInsert, e.EventType, e.CreatedAt, e.Metadata)
	if err != nil {
		log.Fatal("Error while creating new event: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error while creating new event: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	e.ID = id
	return &e, nil
}

func NewEventRepositoryDb(s *sql.DB) EventRepositoryDb {
	return EventRepositoryDb{s}
}
