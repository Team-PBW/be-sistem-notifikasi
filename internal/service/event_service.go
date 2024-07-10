package service

import "golang.org/x/e-calender/internal/repository"

type EventService struct {
	EventRepository *repository.EventRepository
}

// func (e *EventService)