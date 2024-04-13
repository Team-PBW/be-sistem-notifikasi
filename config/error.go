package config

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type HTTPErrorHandler struct {
	StatusCode int
	EchoErr map[string]string
}

func NewCustomErrorHandler() *HTTPErrorHandler {
	return &HTTPErrorHandler{
		EchoErr: make(map[string]string),
	}
}

func GetHTTPError(statusCode int, message []string) *echo.HTTPError {
	return &echo.HTTPError{
		Code: statusCode,
		Message: message,
	}
}

func (httpErr *HTTPErrorHandler) Check(err error, c echo.Context) {
	// _, ok := err.(*echo.HTTPError)
	// if !ok {
	// 	statusCode = http.StatusInternalServerError
	// }


	errorLists, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range errorLists {
			tag := err.Tag()
			switch tag {
            case "required":
                httpErr.EchoErr[tag] = fmt.Sprintf("%s is required", 
                    err.Field())
            case "email":
                httpErr.EchoErr[tag] = fmt.Sprintf("%s is not valid email", 
                    err.Field())
            case "gte":
                httpErr.EchoErr[tag] = fmt.Sprintf("%s value must be greater than %s",
                    err.Field(), err.Param())
            case "lte":
                httpErr.EchoErr[tag] = fmt.Sprintf("%s value must be lower than %s",
                    err.Field(), err.Param())
					
			}
			// break
		}

		var allError []string

		for _, value := range httpErr.EchoErr {
			allError = append(allError, value)
		}
		c.Logger().Error(allError)
		c.JSON(http.StatusBadRequest, allError)
	}
	c.Logger().Error(err.Error())
	c.JSON(http.StatusInternalServerError, err.Error())
}