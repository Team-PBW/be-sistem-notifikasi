package router

import (
	"github.com/labstack/echo"
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

	auth, event := initApp.DefineHandler()

	e := echo.New()
	e.Debug = true

	// middleware
	newJwt := middleware.GetJwtValidate()
	useAuthJwt := newJwt.ValidateJWT
	
	r := e.Group("/api/v1")

	a := r.Group("/auth")
	evt := r.Group("/event")
	// n := r.Group("/notification")

	// auth router
	a.POST("/account", auth.CreateAccount)
	a.POST("/login", auth.Login)

	// event router
	evt.Use(useAuthJwt)
	
	evt.GET("/:id", event.DetailEvent)
	evt.POST("/new", event.AddEvent)

	// notification router

	return e
}