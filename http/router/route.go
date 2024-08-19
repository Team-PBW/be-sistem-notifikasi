package router

import (
	"github.com/labstack/echo"
	midware "github.com/labstack/echo/middleware"
	// "github.com/labstack/echo/v4"
	"golang.org/x/e-calender/app"
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

	e.Use(midware.CORSWithConfig(midware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // Allow your frontend origin
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))


	r := e.Group("/api/v1")

	// auth routes
	authGroup := r.Group("/auth")
	authGroup.POST("/register", auth.CreateAccount)
	authGroup.POST("/login", auth.Login)
	authGroup.GET("/me", auth.GetMe, useAuthJwt)
	authGroup.GET("/dropdown", auth.CheckFollower, useAuthJwt)

	// event routes
	eventGroup := r.Group("/event")
	eventGroup.Use(useAuthJwt)
	eventGroup.GET("/:id", event.DetailEvent)
	eventGroup.POST("", event.AddEvent)
	eventGroup.GET("/all_events", event.CategorizeEventByDatetime)
	// eventGroup.PATCH("/:id", event.UpdateEvent)
	// eventGroup.DELETE("/:id", event.DeleteEvent)

	// category routes
	categoryGroup := r.Group("/category")
	categoryGroup.Use(useAuthJwt)
	categoryGroup.POST("", category.CreateCategory)
	categoryGroup.DELETE("/:id", category.DeleteCategory)
	categoryGroup.GET("/categories", category.FindAllCategory)

	// notification routes
	notificationGroup := r.Group("/notification")
	notificationGroup.Use(useAuthJwt)
	notificationGroup.GET("/fetch", notification.GetAllNotification)
	notificationGroup.GET("/cron", notification.NotificationBroadcast)
	// notificationGroup.POST("", notification.CreateNotification)

	return e
}
