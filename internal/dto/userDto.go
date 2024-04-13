package dto

type UserDto struct {
	Username    string `json:"username"`
	Pass		string `json:"pass"`
	PhoneNumber string `json:"phone_number"`
}