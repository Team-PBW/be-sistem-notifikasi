package entity

import (
	"time"

	// "github.com/google/uuid"
	// "github.com/golang/protobuf/ptypes/timestamp"
)

type EventEntity struct {
	Id          string `gorm:"primary_key"`
	CategoryId  int
	Title       string
	Description string
	StartTime   time.Time `gorm:"type:datetime"`
	EndTime     time.Time `gorm:"type:datetime"`
	// TimeDistance int
	Location string
	// Distance int
	Date time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Bentrok bool
}