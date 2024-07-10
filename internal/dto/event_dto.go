// lokasi
// judul
// tanggal mulai
// jam mulai
// jam selesai

// user yang diundang

package dto

// import "time"

type EventDto struct {
	Id string `json:"id_event"`
	Title string `json:"title"`
	IdCategory int `json:"id_category"`
	Location string `json:"location"`
	// Distance int `json:"distance"` // km
	// TimeDistance int `json:"time_distance"` //hour or min
	Description string `json:"description"`
	Date string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	InvitedUser []string `json:"invited_user"`
}