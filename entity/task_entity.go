package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type TaskEntity struct {
	Id          string
	CategoryId  string
	Title       string
	Description string
	DueDate     *timestamp.Timestamp
	Priority int
	Status string
	CreatedAt *timestamp.Timestamp
	UpdatedAt *timestamp.Timestamp
}