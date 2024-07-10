package entity

import (
	"time"

	// "github.com/golang/protobuf/ptypes/timestamp"
)

type UserEntity struct {
	Username  string  `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time
}