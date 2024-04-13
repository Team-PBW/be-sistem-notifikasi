package app

import (
	"log"

	// "golang.org/x/e-calender/http/handler"
	"github.com/sirupsen/logrus"
	"golang.org/x/e-calender/config"
	"golang.org/x/e-calender/database"
	"golang.org/x/e-calender/http/handler"
	"golang.org/x/e-calender/http/handler/helper"
	"golang.org/x/e-calender/internal/repository"
	"golang.org/x/e-calender/internal/service/auth"
	"golang.org/x/e-calender/internal/service/event"
	"gorm.io/gorm"
)

type InitApp struct {
	DB *gorm.DB
}

func NewInitApp(transaction *gorm.DB) *InitApp {
	return &InitApp{
		DB: transaction,
	}
}

func GetDatabase() *gorm.DB {
	db, err := database.NewGorm()
	// db.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("Failed to initialize database:", err)
		return nil
	}

	tx := db.Begin()
	return tx
}

func (a *InitApp) DefineHandler() (*handler.AuthHandler, *handler.EventHandler) {

	log := logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetLevel(logrus.InfoLevel)

	authRepository := repository.NewAuthRepository(a.DB)
	eventRepository := repository.NewEventRepository(a.DB)
	// notificationRepository := repository.NewNotificationRepository(a.DB)

	authService := auth.NewAuthService(authRepository)
	eventService := event.NewEventService(eventRepository, config.NewValidator())
	// notificationService := notification.NewNotificationService(notificationRepository)

	authHandler := handler.NewAuthHandler(authService, helper.NewHelper(), log)
	eventHandler := handler.NewEventHandler(eventService, helper.NewHelper())
	// notificationHandler := handler.NewNotificationHandler(notificationService)

	return authHandler, eventHandler
}
