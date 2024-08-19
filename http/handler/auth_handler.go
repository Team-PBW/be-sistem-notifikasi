package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"golang.org/x/e-calender/http/handler/helper"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/service/auth"
	"golang.org/x/e-calender/model"
)

type AuthHandler struct {
	Do          *helper.HelperHandler
	AuthService *auth.AuthService
	Log         *logrus.Logger
}

func NewAuthHandler(authService *auth.AuthService, helperHandler *helper.HelperHandler, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Do:          helperHandler,
		Log: log,
	}
}

// func NewAuthHandler() *AuthHandler {
// 	return &AuthHandler{}
// }

func GetStructUser() *model.User {
	return &model.User{}
}

func (u *AuthHandler) CreateAccount(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	if c.Request().Method != "POST" {
		log.Println("Create account only executed with POST Method")
		return echo.NewHTTPError(http.StatusBadRequest, "Create account only executed with POST Method")
	}

	// user := GetStructUser()
	// err := u.Do.DecodeJson(c.Request().Body, user)
	// if err != nil {
	// 	// c.Logger().Error(err)
	// 	// log.Println("Failed to decode JSON")
	// 	u.Log.Errorf("failed to decode json: %s", err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode JSON")
	// }

	var user *dto.UserDto
	d := json.NewDecoder(c.Request().Body)
	if err := d.Decode(&user); err != nil {
		log.Println("Failed to decode JSON")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode JSON")
	}

	// if phoneNumber, err := strconv.Atoi(user.PhoneNumber); err == nil {
	// 	user.PhoneNumber = phoneNumber
	// } else {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Failed to convert phone number to int")
	// }
	

	deadline := time.Now().Add(10 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	err := u.AuthService.CreateAccount(ctx, user)
	if err != nil {
		c.Logger().Error(err)
		// log.Println("Failed to create account")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create account")
	}

	response := &dto.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success create account",
	}

	if err := json.NewEncoder(c.Response()).Encode(response); err != nil {
		c.Logger().Error(err)
		// log.Println("Failed to encode data")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to encode data")
	}

	return nil
}

func (u *AuthHandler) Login(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	if c.Request().Method != "POST" {
		log.Println("Create account only executed with POST Method")
		return echo.NewHTTPError(http.StatusBadRequest, "Create account only executed with POST Method")
	}

	var user *dto.UserDto
	d := json.NewDecoder(c.Request().Body)
	if err := d.Decode(&user); err != nil {
		log.Println("Failed to decode")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode")
	}

	deadline := time.Now().Add(3 * time.Minute)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	token, err := u.AuthService.Find(ctx, user)
	if err != nil {
		switch err {
		case u.AuthService.ErrUserNotFound():
			log.Println("User Not Found")
			return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
		case u.AuthService.ErrInvalidPassword():
			log.Println("Password not match")
			return echo.NewHTTPError(http.StatusUnauthorized, "Password not match")
		default:
			log.Println("Internal server error, cant find data")
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error, cant find data")
		}
	}

	response := &dto.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    token,
	}

	if err := json.NewEncoder(c.Response()).Encode(response); err != nil {
		log.Println("Failed to encode data")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to encode data")
	}

	return nil
	// w.Write([]byte(fmt.Sprintf("access_token: %s", token)))
}

func (u *AuthHandler) GetMe(c echo.Context) error {
	username := c.Get("username").(string)

	log.Println(username)

	getMe, err := u.AuthService.GetUserInformation(username);
	if err != nil {
		c.Logger().Error()
		return err
	}

	

	return c.JSON(200, getMe)
}


func (u *AuthHandler) CheckFollower(c echo.Context) error {
	username := c.Get("username").(string)

	keyword := c.QueryParam("following")

	getFollowingDropdown, err := u.AuthService.BatchUserDropdown(username, keyword);
	if err != nil {
		c.Logger().Error()
		return c.JSON(500, http.StatusInternalServerError)
	}

	if len(getFollowingDropdown) == 0 {
		return echo.NewHTTPError(404, "Follower Empty")
	}

	return c.JSON(200, getFollowingDropdown)
}