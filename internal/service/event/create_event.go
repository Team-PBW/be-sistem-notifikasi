package event

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/e-calender/config"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/repository"

	// "golang.org/x/e-calender/internal/service/event"
	"golang.org/x/e-calender/model"
)

type EventService struct {
	EventRepository *repository.EventRepository
	NotificationRepository *repository.NotificationRepository
	Validate        *config.CustomValidator
}

func NewEventService(e *repository.EventRepository, n *repository.NotificationRepository, v *config.CustomValidator) *EventService {
	return &EventService{
		EventRepository: e,
		NotificationRepository: n,
		Validate:        v,
	}
}

func TimeConvert(format int, sign string, times ...string) ([]int, error) {
	var valToStr []string
	var valToInt []int
	for _,val := range times {
		valToStr = strings.SplitAfter(val, sign)
		evtInt, err := strconv.Atoi(valToStr[format])
		if err != nil {
			return nil, err
		}
		valToInt = append(valToInt, evtInt)
	}
	return valToInt, nil
}

func (e *EventService) CreateEvent(user string, event *dto.EventDto) (interface{}, error) {
	userToModel := &model.User{
		Username: user,
	}

	err := e.Validate.TryValidate(userToModel, event)
	if err != nil {
		return nil, err
	}

	// startTime := strings.SplitAfter(event.StartTime, ":")
	// endTime := strings.SplitAfter(event.EndTime, ":")

	// evtStartInt, err := strconv.Atoi(startTime[0])
	// if err != nil {
	// 	return nil, err
	// }

	// evtEndInt, err := strconv.Atoi(endTime[0])
	// if err != nil {
	// 	return nil, err
	// }
	timeFormat, err := TimeConvert(0, ":", event.StartTime, event.EndTime)
	if err != nil {
		return nil, err
	}
	
	for i := 1; i < len(timeFormat); i++ {
		if timeFormat[i-1] > timeFormat[i] {
			return nil, errors.New("event_time: end time must be later than start time")
		}
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

	startTimeUTCMinus8 := startTime.Add(-8 * time.Hour)
	endTimeUTCMinus8 := endTime.Add(-8 * time.Hour)

	id := uuid.New()
	stringID := id.String()
	newEvent := &entity.EventEntity{
		Id:           stringID,
		Title:        event.Title,
		CategoryId:   event.IdCategory,
		Date:         eventDate,
		// TimeDistance: event.TimeDistance,
		Location:     event.Location,
		// Distance:     event.Distance,
		StartTime:    startTimeUTCMinus8,
		EndTime:      endTimeUTCMinus8,
		CreatedAt:    time.Now(),
		Bentrok: false,
	}

	// person := make(map[string][]string)

	// masukin array person ke followed_event_entity

	err = e.EventRepository.EventFollowedPerson(event.InvitedUser)
	if err != nil {
		return nil, err
	}

	err = e.EventRepository.CreateEvent(user, newEvent, event.InvitedUser)
	if err != nil {
		return nil, err
	}

	log.Println(endTimeUTCMinus8)

	yesnt := e.EventRepository.CheckEventExist(startTime, endTime, eventDate)
	if yesnt == true {
		// return nil, nil

		// jalanin create notifikasi warning sudah ada event di jam itu\

		id := uuid.New()
		notifUUID := id.String()

		err := e.EventRepository.UpdateBentrok(stringID)
		if err != nil {
			return nil, err
		}

		notificationWarn := &entity.EventNotification{
			Id: notifUUID, 
			EventId: stringID,
			NotificationTime: time.Now(),
			Message: "Jadwal Bentrok, silahkan periksa kembali",
			SendStatus: false,
		}

		err = e.NotificationRepository.CreateWarningEvent(notificationWarn)
		if err != nil {
			return nil, err
		}

		// return jangan buat data
		// return "Periksa", nil
	}

	log.Println(yesnt)

	// err = e.NotificationRepository.Create()

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
