package service

import (
	// "log"
	// "math/rand"
	// "time"

	// "golang.org/x/e-calender/entity"
	// "golang.org/x/e-calender/internal/dto"
	// "log"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/repository"
)

type NotificationService struct {
	NotificationRepository *repository.NotificationRepository
}

func NewNotificationService(notificationRepository *repository.NotificationRepository) *NotificationService {
	return &NotificationService{
		NotificationRepository: notificationRepository,
	}
}

func (n *NotificationService) CreateNotification(username string) error {
	idEvent, err := n.NotificationRepository.CheckAndFetchId(username)
	if err != nil {
		return err
	}

	flag, startTime, err := n.NotificationRepository.CheckDate(idEvent)
	if err != nil || !flag {
		return err
	}

	hMinusOne := startTime.AddDate(0, 0, -1)

	if time.Now().After(hMinusOne) {
		notificationExists, existingNotifications, err := n.NotificationRepository.CheckNotificationExists(idEvent, hMinusOne)
		if err != nil {
			return err
		}

		if notificationExists {
			log.Println("Existing notifications for the event:", existingNotifications)
			return nil
		}

		id := uuid.New().String()

		eventNotificationData := &entity.EventNotification{
			Id:               id,
			EventId:          idEvent,
			NotificationTime: time.Now(),
			Message:          "Acara akan diadakan H-1",
			SendStatus:       false,
		}

		err = n.NotificationRepository.Create(eventNotificationData)
		if err != nil {
			return err
		}
	}

	return nil
}