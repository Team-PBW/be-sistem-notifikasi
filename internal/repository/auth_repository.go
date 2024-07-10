package repository

import (
	// "errors"
	// "net/http"
	"context"
	"errors"
	"log"

	"golang.org/x/e-calender/entity"
	// "golang.org/x/e-calender/model"
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type AuthRepository struct {
	DB *gorm.DB
}

// type PostRepositoryIntrfc interface {
// 	Create() (*model.User, interface{}, error)
// 	Delete() error
// 	Update(username string) (interface{}, error)
// }

var (
	ErrUserNotFound = errors.New("can't find account")
)

func NewAuthRepository(tx *gorm.DB) *AuthRepository {
	log.Println("user repository")
	return &AuthRepository{
		DB: tx,
	}
}

func (u *AuthRepository) BeginTransaction() *gorm.DB {
	return u.DB.Begin()
}

func (u *AuthRepository) Create(ctx context.Context, user *entity.UserEntity) error {
	log.Println("user")
	defer func() {
		if r := recover(); r != nil {
			u.DB.Rollback()
		}
	}()

	if err := u.DB.Create(&user).Error; err != nil {
		u.DB.Rollback()
		return err
	}

	if err := u.DB.Commit().Error; err != nil{
		return err
	}

	return nil 
}

func (u *AuthRepository) FindAcc(ctx context.Context, username string) (*entity.UserEntity, error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		u.DB.Rollback()
	// 	}
	// }()

	log.Println("username")

	user := &entity.UserEntity{}

	res := u.DB.Where("username = ?", username).Find(user)
	if res.Error != nil {
		log.Println("Error finding user:", res.Error)
		// u.DB.Rollback()
		return nil, res.Error
	}

	// if res.RowsAffected == 0 {
	// 	return nil, ErrUserNotFound
	// }

	return user, res.Error
}

func (u *AuthRepository) GetSelfInformation(username string) (*entity.UserEntity, error) {
	var user *entity.UserEntity
	err := u.DB.Where("email = ?", username).Find(&user).Error

	return user, err
}
