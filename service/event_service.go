package service

import (
	"main/domain"
	"main/dto"
)

type EventService interface {
	GetAllEvents() ([]domain.Event, error)
	CreateEvent(event dto.NewEventRequest) (*dto.NewEventResponse, error)
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (d DefaultEventService) GetAllEvents() ([]domain.Event, error) {
	return d.repo.GetAllEvents()
}

func (d DefaultEventService) CreateEvent(req dto.NewEventRequest) (*dto.NewEventResponse, error) {
	e := domain.Event{
		EventType: req.EventType,
		CreatedAt: req.CreatedAt,
		Metadata:  req.Metadata,
	}
	event, err := d.repo.CreateEvent(e)
	if err != nil {
		return nil, err
	}
	response := event.ToNewEventResponseDto()

	return &response, nil
}

func NewEventService(r domain.EventRepository) DefaultEventService {
	return DefaultEventService{repo: r}
}
