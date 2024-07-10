package repository

import (
	// "github.com/docker/distribution/uuid"
	"log"

	"golang.org/x/e-calender/entity"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	TX *gorm.DB
}

func NewCategoryRepository(tx *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		TX: tx,
	}
}

func (c *CategoryRepository) BeginTransaction() *gorm.DB {
	return c.TX.Begin()
}

func (c *CategoryRepository) Create(data *entity.CategoryEntity, tx *gorm.DB) error {
    log.Println("Starting Create operation")

    if err := tx.Create(&data).Error; err != nil {
        log.Printf("Failed to insert data: %v", err)
        return err
    }

    log.Println("Insert operation successful")
    return nil
}

// func (c *CategoryRepository) FindOrFail(id int, categoryName string) error {
// 	return c.TX.Find(categoryName).Error
// }

func (c *CategoryRepository) Delete( idCategory int ) error {
	var category *entity.CategoryEntity
	return c.TX.Delete(&category, idCategory).Error
}

func (c *CategoryRepository) Update(idCategory int, body interface{}) (*entity.CategoryEntity, error) {
	var category entity.CategoryEntity
	err := c.TX.Model(&category).Where("id = ?", idCategory).Updates(body).Error

	if err != nil {
		return nil, err
	}

	err = c.TX.First(&category, idCategory).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepository) AllCategory(username string, category entity.CategoryEntity) ([]entity.CategoryEntity, error) {
	var categories []entity.CategoryEntity
	err := c.TX.Model(category).Find(&categories).Error

	if err != nil {
		return nil, err
	}
	return categories, nil
}

// func (c *CategoryRepository) Read(username string) ([]entity.CategoryEntity, error) {

// }