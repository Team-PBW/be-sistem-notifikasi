package entity

import (
	"time"
)

type FollowedEventEntity struct {
	EventId    string `gorm:"primaryKey"`
	Username   string `gorm:"foreignKey:Username;references:UserEntity"`
	FollowedAt time.Time
	Confirmed bool
}