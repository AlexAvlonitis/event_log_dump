package domain

import (
	"main/dto"
	"main/errs"
)

type Event struct {
	ID        int64
	EventType string
	CreatedAt string
	Metadata  string
}

type EventRepository interface {
	GetAllEvents() ([]Event, *errs.AppError)
	CreateEvent(Event) (*Event, *errs.AppError)
}

func (e Event) ToNewEventResponseDto() dto.NewEventResponse {
	return dto.NewEventResponse{ID: e.ID}
}
