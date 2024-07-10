package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type TaskNotificationEntity struct {
	Id               string
	TaskId          string
	NotificationTime *timestamp.Timestamp
	Message string
	SendStatus bool
}