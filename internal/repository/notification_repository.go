package repository

import (
	"log"
	"time"

	"github.com/docker/distribution/uuid"
	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/model"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	TX *gorm.DB
}

func NewNotificationRepository(tx *gorm.DB) *NotificationRepository {
	log.Println("notification repository")
	return &NotificationRepository{
		TX: tx,
	}
}

func (n *NotificationRepository) FetchAllNotificationsBeforeDate(beforeDate time.Time) ([]entity.EventNotification, error) {
	var notifications []entity.EventNotification
	err := n.TX.Model(&entity.EventNotification{}).
		Where("notification_time < ?", beforeDate).
		Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *NotificationRepository) CheckNotificationExists(eventID string, beforeDate time.Time) (bool, []entity.EventNotification, error) {
	notifications, err := n.FetchAllNotificationsBeforeDate(beforeDate)
	if err != nil {
		return false, nil, err
	}

	var eventNotifications []entity.EventNotification
	for _, notification := range notifications {
		if notification.EventId == eventID {
			eventNotifications = append(eventNotifications, notification)
		}
	}

	return len(eventNotifications) > 0, eventNotifications, nil
}

func (n *NotificationRepository) CheckAndFetchId(username string) (string, error) {
	var uuid string

	// n.TX.Model(&entity.EventNotification{}).Where("event_id", &uuid).Where(startPoint).First()
	if err := n.TX.Model(&entity.FollowedEventEntity{}).Where("username = ?", username).Pluck("event_id", &uuid).Error; err != nil {
		return "", err
	}
	return uuid, nil
}

func (n *NotificationRepository) CheckDate(idEvent string) (bool, time.Time, error) {
	var startTimeStr string
	err := n.TX.Model(&entity.EventEntity{}).Where("id = ?", idEvent).Select("start_time").Scan(&startTimeStr).Error
	if err != nil {
		return false, time.Time{}, err
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		return false, time.Time{}, err
	}

	return true, startTime, nil
}

func (n *NotificationRepository) Create(event *entity.EventNotification) error {
	
	return n.TX.Model(&entity.EventNotification{}).Create(&event).Error
}

func (n *NotificationRepository) ReadNotification(email string) ([]*entity.EventNotification, error) {
	var notif []*entity.EventNotification

	err := n.TX.Model(&entity.EventNotification{}).Find(&notif).Where("email = ?", email).Order("notification_time DESC").Error
	if err != nil {
		return nil, err
	}
	return notif, nil
}

func (n *NotificationRepository) NotifyUser(id uuid.UUID) ([]*model.User, error) {
	var user []*model.User
	err := n.TX.Model(&model.User{}).Find(&user, "id", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
} 

func (n *NotificationRepository) DeleteNotification(id uuid.UUID) error {
	
	return n.TX.Where("id = ?", id).Delete(&entity.EventNotification{}).Error 
}