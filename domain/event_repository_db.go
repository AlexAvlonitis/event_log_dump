package domain

import (
	"database/sql"
	"log"
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
		err := rows.Scan(&e.LogType, &e.CreatedAt)
		if err != nil {
			log.Fatal("Error while scanning events" + err.Error())
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func NewEventRepositoryDb() EventRepositoryDb {
	client, err := sql.Open("sqlite3", "sqlite3://event-log-dump")
	if err != nil {
		panic(err)
	}

	return EventRepositoryDb{client}
}
