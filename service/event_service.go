package service

import (
	"main/domain"
	"main/dto"
	"main/errs"
)

type EventService interface {
	GetAllEvents() ([]domain.Event, *errs.AppError)
	CreateEvent(event dto.NewEventRequest) (*dto.NewEventResponse, *errs.AppError)
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (d DefaultEventService) GetAllEvents() ([]domain.Event, *errs.AppError) {
	return d.repo.GetAllEvents()
}

func (d DefaultEventService) CreateEvent(req dto.NewEventRequest) (*dto.NewEventResponse, *errs.AppError) {
	e := domain.Event{
		EventType: req.EventType,
		CreatedAt: req.CreatedAt,
		Metadata:  req.Metadata,
	}
	event, error := d.repo.CreateEvent(e)
	if error != nil {
		return nil, error
	}
	response := event.ToNewEventResponseDto()

	return &response, nil
}

func NewEventService(r domain.EventRepository) DefaultEventService {
	return DefaultEventService{repo: r}
}
