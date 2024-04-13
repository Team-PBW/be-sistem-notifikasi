package main

import (
	// "github.com/labstack/echo"
	// "golang.org/x/e-calender/config"
	"golang.org/x/e-calender/http/router"
)

func main() {
	// initService()
	// r := echo.New()

	// router
	routers := router.NewCalenderRouter()
	r := routers.GetAllRouter()

	// custom error handling for validate data
	// errorHandler := config.NewCustomErrorHandler()
	// r.HTTPErrorHandler = errorHandler.Check


	if err := r.Start(":3000"); err != nil {
		r.Logger.Fatal(err)
	}
}

// func initService() {
// 	db, err := database.NewGorm()
// 	if err != nil {
// 		log.Println("cant instatiate db")
// 	}

// 	tx := db.Begin()

// 	getDb, err := db.DB()
// 	if err != nil {
// 		return
// 	}

// 	config.RunPostgresMigrate(getDb)
// 	cv := config.NewValidator()
// 	repo := repository.NewEventRepository(tx)
// 	e := service.NewEventService(repo, cv.Validate)
// 	handler.NewEventHandler(e)
// }

// func Route() {

// }
