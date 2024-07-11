package dto

// import "time"

type EventNotificationDto struct {
	Id string `json:"id_notification"`
	EventId string `json:"id_event"`
	NotificationTime string `json:"notification_time"`
	Title string `json:"title"`
	CategoryId int `json:"id_category"`
	Location string `json:"location"`
	Message string `json:"message"`
	// Distance int `json:"distance"` // km
	// TimeDistance int `json:"time_distance"` //hour or min
	// Description string `json:"description"`
	Date string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Bentrok bool `json:"bentrok"`
	SendStatus bool `json:"send_status"`
}