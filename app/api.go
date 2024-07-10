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
	"golang.org/x/e-calender/internal/service"
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

	// tx := db.Begin()
	return db
}

func (a *InitApp) DefineHandler() (*handler.AuthHandler, *handler.EventHandler, *handler.CategoryHandler, *handler.NotificationHandler) {

	log := logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetLevel(logrus.InfoLevel)

	authRepository := repository.NewAuthRepository(a.DB)
	eventRepository := repository.NewEventRepository(a.DB)
	categoryRepository := repository.NewCategoryRepository(a.DB)
	notificationRepository := repository.NewNotificationRepository(a.DB)

	authService := auth.NewAuthService(authRepository)
	eventService := event.NewEventService(eventRepository, config.NewValidator())
	categoryService := service.NewCategoryService(categoryRepository)
	notificationService := service.NewNotificationService(notificationRepository)

	authHandler := handler.NewAuthHandler(authService, helper.NewHelper(), log)
	eventHandler := handler.NewEventHandler(eventService, helper.NewHelper())
	categoryHandler := handler.NewCategoryHandler(categoryService)
	notificationHandler := handler.NewNotificationHandler(notificationService, helper.NewHelper())

	log.Println(categoryHandler)
	// notificationHandler := handler.NewNotificationHandler(notificationService)

	return authHandler, eventHandler, categoryHandler, notificationHandler
}
