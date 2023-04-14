package domain

import "time"

type EventRepositoryStub struct {
	events []Event
}

func (er EventRepositoryStub) GetAllEvents() ([]Event, error) {
	return er.events, nil
}

func NewEventRepositoryStub() EventRepositoryStub {
	events := []Event{
		{LogType: "UserCreated", CreatedAt: time.Now()},
		{LogType: "UserUpdated", CreatedAt: time.Now()},
	}

	return EventRepositoryStub{events}
}
