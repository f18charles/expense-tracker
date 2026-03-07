package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepo struct {}

func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{}
}

func (cr *CategoryRepo) CreateCategory(cat *models.Category) error {
	return database.DB.Create(cat).Error
}

func (cr *CategoryRepo) GetCategoryByID(id uuid.UUID) (*models.Category, error) {
	var cat models.Category
	result := database.DB.Where("id = ?",id).First(&cat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &cat, nil
}

func (cr *CategoryRepo) UpdateCategory(cat *models.Category) error {
	return database.DB.Save(cat).Error
}

func (cr *CategoryRepo) ListCategory(user_id uuid.UUID) ([]models.Category, error) {
	cats := []models.Category{}
	results := database.DB.Where("user_id = ?", user_id).Find(&cats)
	if results.Error != nil {
		return nil,results.Error
	}
	return cats, nil
}

func (cr *CategoryRepo) DeleteCategory(id uuid.UUID) error {
	result := database.DB.Delete(&models.Category{},"id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
