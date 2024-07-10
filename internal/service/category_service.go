package service

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/e-calender/entity"
	"golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/internal/repository"
)

type CategoryService struct {
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryService(categoryRepository *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryService) CreateService(category *dto.CategoryDto, username string) error {
	rand.Seed(time.Now().UnixNano())

    // Generate a random Id between 100000 and 999999
    randomId := rand.Intn(900000) + 100000

	data := &entity.CategoryEntity{
		Id: randomId,
		Name: category.Name,
		Description: category.Description,
		Username: username,
	}
	 // Start a new transaction
	 tx := c.CategoryRepository.BeginTransaction()
	 if tx.Error != nil {
		 return tx.Error
	 }
 
	 // Defer rollback in case of panic
	 defer func() {
		 if r := recover(); r != nil {
			 tx.Rollback()
		 }
	 }()

	 if err := c.CategoryRepository.Create(data, tx); err != nil {
        tx.Rollback()
        return err
    }

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) Delete(idCategory int) error {
	err := c.CategoryRepository.Delete(idCategory)

	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) GetCategory(username string) (error, []entity.CategoryEntity) {
	var category entity.CategoryEntity

	data, err := c.CategoryRepository.AllCategory(username, category)

	log.Println(data)

	if err != nil {
		return err, nil
	}

	return nil, data
}