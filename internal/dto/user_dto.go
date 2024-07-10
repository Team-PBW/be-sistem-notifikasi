package dto

type UserDto struct {
	Username    string `json:"username"`
	Email	string `json:"email"`
	Password		string `json:"password"`
	Alamat	string `json:"alamat"`
	PhoneNumber string `json:"phone_number"`
}