package router

import (
	"github.com/labstack/echo"
	"golang.org/x/e-calender/app"
	// "golang.org/x/e-calender/internal/service/notification"
	"golang.org/x/e-calender/middleware"
)

type CalenderRouter struct{}

func NewCalenderRouter() *CalenderRouter {
	return &CalenderRouter{}
}

func (c *CalenderRouter) GetAllRouter() *echo.Echo {
	db := app.GetDatabase()
	initApp := app.NewInitApp(db)

	auth, event, category, notification := initApp.DefineHandler()

	e := echo.New()
	e.Debug = true

	// middleware
	newJwt := middleware.GetJwtValidate()
	useAuthJwt := newJwt.ValidateJWT
	
	r := e.Group("/api/v1")

	a := r.Group("/auth")
	evt := r.Group("/event")
	cat := r.Group("/category")
	notif := r.Group("/notification")
	// info := r.Group("/info")

	// n := r.Group("/notification")

	// auth router
	a.POST("/register", auth.CreateAccount)
	a.POST("/login", auth.Login)

	// evt.Use(useAuthJwt)
	a.GET("/me", auth.GetMe, useAuthJwt)

	// event router
	evt.Use(useAuthJwt)
	notif.Use(useAuthJwt)
	
	evt.GET("/:id", event.DetailEvent)
	evt.POST("/event", event.AddEvent)
	evt.GET("/events", event.CategorizeEventByDatetime)
	// evt.PATCH("/:id", event.UpdateEvent)
	// evt.DELETE("/:id", event.DeleteEvent)

	cat.Use(useAuthJwt)
	cat.POST("/category", category.CreateCategory)
	cat.DELETE("/category/:id", category.DeleteCategory)
	cat.GET("/categories", category.FindAllCategory)

	notif.GET("/notif", notification.NotificationBroadcast)
	// notification.GET("", notification.CreateNotification)


	// notification router


	return e
}