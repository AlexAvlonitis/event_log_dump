package domain

import "main/errs"

type EventRepositoryStub struct {
	events []Event
}

func (er EventRepositoryStub) GetAllEvents() ([]Event, *errs.AppError) {
	return er.events, nil
}

func NewEventRepositoryStub() EventRepositoryStub {
	events := []Event{
		{EventType: "UserCreated", CreatedAt: "2022-02-02"},
		{EventType: "UserUpdated", CreatedAt: "2022-03-03"},
	}

	return EventRepositoryStub{events}
}
