package entity

import (
	"time"

	// "github.com/golang/protobuf/ptypes/timestamp"
)

type EventNotification struct {
	Id               string `gorm:"primaryKey"`
	EventId          string `gorm:"foreignKey:Id;references:EventEntity"`
	NotificationTime time.Time
	Message string
	SendStatus bool
}