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
	"net/smtp"

	"github.com/google/uuid"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/dto"
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

func (n *NotificationService) NotificationAlarm(data string) error {
	// cek nama dan hari h notifikasi
	

	// var guests []string	

	// kirim ke email

	// Configuration
	from := "tachibanahinata2021@gmail.com"
	password := "wdql sori qybe hksj"
	to := []string{"gunkwibawa17@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	
	message := []byte("Your account has been hacked")
	
	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)
	
	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	// buatkan notif
	return nil
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

	log.Println(startTime)

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

	// hMinusOne := startTime.AddDate(0, 0, -1)

	// if time.Now().After(hMinusOne) {
	// 	notificationExists, existingNotifications, err := n.NotificationRepository.CheckNotificationExists(idEvent, hMinusOne)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if notificationExists {
	// 		log.Println("Existing notifications for the event:", existingNotifications)
	// 		return nil
	// 	}
	// }

	return nil
}

func (n *NotificationService) GetAllNotification(username string) ([]*dto.EventNotificationDto, error) {
	var allEvents []*dto.EventNotificationDto

	notifs, err := n.NotificationRepository.ReadNotification(username)
	if err != nil {
		return nil, err
	}

	for _, event := range notifs {
		// date := event.Date.Format("2006-05-01")
		// startTime := event.StartTime.Format("08:01:09")
		// endTime := event.EndTime.Format("08:01:09")
		newEvent := &dto.EventNotificationDto {
			Id: event.Id,
			EventId: event.EventId,
			NotificationTime: event.NotificationTime.Format("2006-01-02"),
			Title: event.Title,
			CategoryId: event.CategoryId,
			Location: event.Location,
			Message: event.Message,
			// Description: event.Description,
			Date:        event.Date.Format("2006-01-02"),      // Format YYYY-MM-DD
			StartTime:   event.StartTime.Format("15:04:05"),   // Format HH:MM:SS
			EndTime:     event.EndTime.Format("15:04:05"),     // Format HH:MM:SS
			Bentrok: event.Bentrok,
			SendStatus: event.SendStatus,
		}

		allEvents = append(allEvents, newEvent)
	}

	return allEvents, nil
}

func (n* NotificationService) CheckDateEvent(username string) error {
	return nil
}