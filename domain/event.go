package domain

import (
	"main/dto"
)

type Event struct {
	ID        int64
	EventType string
	CreatedAt string
	Metadata  string
}

type EventRepository interface {
	GetAllEvents() ([]Event, error)
	CreateEvent(Event) (*Event, error)
}

func (e Event) ToNewEventResponseDto() dto.NewEventResponse {
	return dto.NewEventResponse{ID: e.ID}
}
