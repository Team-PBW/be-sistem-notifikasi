package entity

import (
	"time"

	// "github.com/golang/protobuf/ptypes/timestamp"
)

type EventJoinNotificationEntity struct {
	Id               string `gorm:"primaryKey"`
	EventId          string `gorm:"foreignKey:Id;references:EventEntity"`
	NotificationTime time.Time
	Location string
	Message string
	SendStatus bool
	CategoryId int
	Title string
	Date time.Time
	StartTime time.Time
	EndTime time.Time
	Bentrok bool
}