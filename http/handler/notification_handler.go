package handler

import (
	// "time"

	"net/http"
	"github.com/labstack/echo"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"golang.org/x/e-calender/http/handler/helper"
	"golang.org/x/e-calender/internal/service"
)

type NotificationHandler struct {
	NotificationService *service.NotificationService
	Do           *helper.HelperHandler
}

func NewNotificationHandler(notifService *service.NotificationService, helperHandler *helper.HelperHandler) *NotificationHandler {
	return &NotificationHandler{
		Do:           helperHandler,
		NotificationService: notifService,
	}
}

func (n *NotificationHandler) NotificationBroadcast(c echo.Context) error {
	log.Info("Create new cron")
	cronNew := cron.New()

	username := c.Get("username").(string)	

    // Menjadwalkan job yang akan berjalan setiap satu hari
    cronNew.AddFunc("0 * * * * *", func() {
        log.Info("[Job] job every minute")
        // Tambahkan logika pekerjaan yang Anda inginkan di sini

		// cek notifikasi dari nama wawan dimana cron_saat_ini h-25jam start_date (join followed_event dgn event_entities)
		// n.NotificationService.Check
		n.NotificationService.NotificationAlarm(username)
    })

	log.Info("Start cron")
	cronNew.Start()
	// printCronEntries(cronNew.Entries())
	// time.Sleep(2 * time.Second)

	return c.JSON(http.StatusOK, nil)
}

func (n *NotificationHandler) GetAllNotification(c echo.Context) error {
	username := c.Get("username").(string)

	datas, err := n.NotificationService.GetAllNotification(username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, datas)
}