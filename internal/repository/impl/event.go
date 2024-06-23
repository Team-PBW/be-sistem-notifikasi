package impl

import (
	"golang.org/x/e-calender/model"
	"github.com/google/uuid"
)

type DBTX interface {
	CreateEvent() error
	UpdateEvent(id string) (*model.Event, error)
	FindGuestsInEvent(id string) ([]string, error)
	UpdateGuestByEventID(id string, guests []*model.EveryPerson) (*model.EventPersonConfirmed, error)
	FindEventByID(id string) (*model.Event, error)
	FindEventsByHost(username string) ([]*model.Event, error)
	DeleteEvent(id string) error
}

type NotificationDBTX interface {
	NotifyUser(email string) error
	ReadNotification(email string) ([]*model.Notification, error)
	DeleteNotification(id uuid.UUID) error
}