package entity

type CategoryEntity struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Username    string `gorm:"column:username"`
}
