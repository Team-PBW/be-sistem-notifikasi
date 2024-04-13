package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"golang.org/x/e-calender/http/handler/helper"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/service/event"
)

var (
	BASE_URL = "http://localhost:8000/user"
)

var (
	internalServerError = 500
)

type EventHandler struct {
	*http.Client
	EventService *event.EventService
	Do           *helper.HelperHandler
}

func NewEventHandler(eventService *event.EventService, helperHandler *helper.HelperHandler) *EventHandler {
	return &EventHandler{
		Do:  helperHandler,
		Client:       new(http.Client),
		EventService: eventService,
	}
}

func (e *EventHandler) AddEvent(c echo.Context) error {

	log.Println("this is add event")

	if c.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	// // nembak ke user service
	// req, err := http.NewRequest("GET", BASE_URL, nil)
	// if err != nil {
	// 	// http.Error(c.Writer, "400", http.StatusBadRequest)
	// 	return echo.NewHTTPError(http.StatusBadRequest, "must get method")
	// }

	// resp, err := e.Client.Do(req)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	// }
	// defer resp.Body.Close()

	// var userDto *dto.Response

	// err = json.NewDecoder(resp.Body).Decode(&userDto.Data)
	// if err != nil {
	// 	// http.Error(c.Writer, "404", http.StatusNotFound)
	// 	// echo.NewHTTPError(http.StatusNotFound, gin.H{"status": "Data not found"})
	// 	// return
	// 	responseUser := &dto.Response{
	// 		StatusCode: 404,
	// 		Status:     "404 Not Found",
	// 		Data:       "Data not found",
	// 	}
	// 	return echo.NewHTTPError(responseUser.StatusCode, responseUser.Data)
	// }

	// get user request from addevent contoller

	var event *dto.EventDTO

	username := c.Get("username")

	log.Println(username)

	// err := json.NewDecoder(c.Request().Body).Decode(&event)
	err := e.Do.DecodeJson(c.Request().Body, &event)
	if err != nil {
		log.Println("error decode event")
		return echo.NewHTTPError(500, "error internal server")
	}

	// decode to struct in json
	// if err = c.ShouldBindJSON(&dataCreated); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	// }

	responseUser := &dto.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Data:       username,
	}

	if convertUsername, ok := responseUser.Data.(string); ok {
		data, err := e.EventService.CreateEvent(convertUsername, event)
		if err != nil {
			log.Println("error create event")
			return echo.NewHTTPError(internalServerError, "internal service error")
		}

		if data == nil {
			log.Println("error data not found")
			return echo.NewHTTPError(http.StatusNotFound, "error 404")
		}
	}

	return c.JSON(http.StatusOK, "200 success create event")
}

// func (e *EventHandler) UpdateEvent(c echo.Context) error {

// 	if c.Request().Method != http.MethodPatch {
// 		return echo.NewHTTPError(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
// 	}

// 	var evt *dto.EventDTO
// 	err := json.NewDecoder(c.Request().Body).Decode(&evt)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
// 	}

// 	id := c.Param("id")

// 	eventObj, err := e.EventService.UpdateEvent(id, evt)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
// 	}

// 	return c.JSON(http.StatusOK, dto)
// }

func (e *EventHandler) DeleteEvent(c echo.Context) error {
	if c.Request().Method != http.MethodDelete {
		return echo.NewHTTPError(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
	}

	id := c.Param("id")

	err := e.EventService.DeleteEvent(id)
	if err != nil {
		return err
		
	}

	return c.JSON(http.StatusOK, gin.H{"status": "event deleted"})
}

func (e *EventHandler) DetailEvent(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
	}

	id := c.Param("id")

	evt, err := e.EventService.DetailEvent(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, gin.H{"error": "500 Internal Server Error"})
	}

	res := &dto.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Data:       evt,
	}

	return c.JSON(http.StatusOK, res)
}

func (e *EventHandler) ShowAllEvents(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusBadRequest, gin.H{"error": "400 Bad Request"})
	}

	return c.JSON(http.StatusOK, gin.H{"status": "event deleted"})
}
