package domain

import (
	"time"
)

type Event struct {
	LogType   string    `json:"log_type"`
	CreatedAt time.Time `json:"created_at"`
}

type EventRepository interface {
	GetAllEvents() ([]Event, error)
}
