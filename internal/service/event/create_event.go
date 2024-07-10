package event

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/e-calender/config"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/repository"
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

func (e *EventService) CreateEvent(user string, event *dto.EventDto) (interface{}, error) {
	userToModel := &model.User{
		Username: user,
	}

	err := e.Validate.TryValidate(userToModel, event)
	if err != nil {
		return nil, err
	}

	const (
		dateLayout     = "2006-01-02"
		timeLayout     = "15:04:05"
		dateTimeLayout = "2006-01-02 15:04:05"
	)

	eventDate, err := time.Parse(dateLayout, event.Date)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil, err
	}

	startDateTimeStr := event.Date + " " + event.StartTime
	startTime, err := time.Parse(dateTimeLayout, startDateTimeStr)
	if err != nil {
		log.Println("Error parsing start time:", err)
		return nil, err
	}

	endDateTimeStr := event.Date + " " + event.EndTime
	endTime, err := time.Parse(dateTimeLayout, endDateTimeStr)
	if err != nil {
		log.Println("Error parsing end time:", err)
		return nil, err
	}

	id := uuid.New()
	stringID := id.String()
	newEvent := &entity.EventEntity{
		Id:           stringID,
		Title:        event.Title,
		CategoryId:   event.IdCategory,
		Date:         eventDate,
		TimeDistance: event.TimeDistance,
		Location:     event.Location,
		Distance:     event.Distance,
		StartTime:    startTime,
		EndTime:      endTime,
		CreatedAt:    time.Now(),
	}

	person := make(map[string][]string)

	err = e.EventRepository.CreateEvent(user, newEvent, person)
	if err != nil {
		return nil, err
	}

	users, err := e.EventRepository.CheckPersonExist(stringID, event.InvitedUser)
	if err != nil {
		log.Println("Error checking person existence:", err)
		return nil, err
	}

	evtFollow := make([]*entity.FollowedEventEntity, 0)

	for _, val := range users {
		evtFollowed := &entity.FollowedEventEntity{
			EventId:   stringID,
			Username:  val.Username,
			Confirmed: false,
		}

		evtFollow = append(evtFollow, evtFollowed)
	}

	err = e.EventRepository.InvitePersonToEvent(evtFollow)
	if err != nil {
		log.Println("Error inviting person to event:", err)
		return nil, err
	}

	return event, nil
}
