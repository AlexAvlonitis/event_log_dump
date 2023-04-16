package dto

import (
	"time"
)

type NewEventRequest struct {
	EventType string `json:"event_type"`
	CreatedAt string `json:"created_at"`
	Metadata  string `json:"metadata"`
}

func (r NewEventRequest) Validate() error {
	_, error := time.Parse(r.CreatedAt, r.CreatedAt)
	if error != nil {
		return error
	}

	return nil
}
