package service

import "main/domain"

type EventService interface {
	GetAllEvents() ([]domain.Event, error)
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (d DefaultEventService) GetAllEvents() ([]domain.Event, error) {
	return d.repo.GetAllEvents()
}

func NewEventService(r domain.EventRepository) DefaultEventService {
	return DefaultEventService{repo: r}
}
