package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/service"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
	}
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var category dto.CategoryDto

	// data := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&category)

	if err != nil {
		c.Logger().Error()
		return err
	}

	username := c.Get("username").(string)

	err = h.CategoryService.CreateService(&category, username)
	if err != nil {
		c.Logger().Error()
		return err
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Category created successfully",
	})
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	var idCategory string
	idCategory = c.Param("id")

	id, err := strconv.Atoi(idCategory)
	// err := json.NewDecoder(c.Request().Body).Decode(&idCategory)

	if err != nil {
		c.Logger().Error()
		return err
	}

	err = h.CategoryService.Delete(id)

	if err != nil {
		c.Logger().Error()
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Category deleted",
	})
}

func (h *CategoryHandler) FindAllCategory(c echo.Context) error {
	// var categories []dto.CategoryDto
	username := c.Get("username").(string)

	err, categories := h.CategoryService.GetCategory(username)

	if err != nil {
		c.Logger().Error()
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": categories,
	})
}
