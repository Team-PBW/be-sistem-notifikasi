package event

import (
	"golang.org/x/e-calender/config"
	"golang.org/x/e-calender/internal/repository"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/model"
)

type EventService struct {
	EventRepository *repository.EventRepository
	Validate        *config.CustomValidator
}

func NewEventService(e *repository.EventRepository, v *config.CustomValidator) *EventService {
	return &EventService{
		EventRepository: e,
		Validate:        v,
	}
}

func (e *EventService) CreateEvent(user string, event *dto.EventDTO) (interface{}, error) {

	userToModel := &model.User{
		Username: user,
	}

	err := e.Validate.TryValidate(userToModel, event)
	if err != nil {
		return nil, err
	}

	var id string
	// generateID()

	newEvent := &model.Event{
		// Id: id,
		EventName:     event.EventName,
		FromDate:      event.FromDate,
		ToDate:        event.ToDate,
		EventLocation: event.Location,
		Descriptions:  event.Description,
	}

	person := make(map[string][]string)

	var everyPerson []string
	for _, val := range event.PersonConfirmed {
		everyPerson = append(everyPerson, val.Username)
	}
	person[id] = everyPerson

	err = e.EventRepository.CreateEvent(user, newEvent, person)
	if err != nil {
		return nil, err
	}

	// event := &dto.EventDTO{
	// 	ID:
	// }

	return event, nil
}
