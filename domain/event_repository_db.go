package domain

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type EventRepositoryDb struct {
	client *sql.DB
}

func (er EventRepositoryDb) GetAllEvents() ([]Event, error) {
	findAllSql := "select * from events"

	rows, err := er.client.Query(findAllSql)
	if err != nil {
		log.Fatal("Error while querying events table" + err.Error())
		return nil, err
	}

	events := make([]Event, 0)
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.EventType, &e.CreatedAt, &e.Metadata)
		if err != nil {
			log.Fatal("Error while scanning events" + err.Error())
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func (er EventRepositoryDb) CreateEvent(e Event) (*Event, error) {
	sqlInsert := "INSERT INTO events (event_type, created_at, metadata) values (?, ?, ?)"

	result, err := er.client.Exec(sqlInsert, e.EventType, e.CreatedAt, e.Metadata)
	if err != nil {
		log.Fatal("Error while creating new event: " + err.Error())
		return nil, errors.New("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error while creating new event: " + err.Error())
		return nil, errors.New("Unexpected error from database")
	}
	e.ID = id
	return &e, nil
}

func NewEventRepositoryDb(s *sql.DB) EventRepositoryDb {
	return EventRepositoryDb{s}
}
