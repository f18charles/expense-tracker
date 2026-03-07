package services

import (
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	category_repo *repository.CategoryRepo
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		category_repo: repository.NewCategoryRepo(),
	}
}

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Color string `json:"color" binding:"required"`
	Icon string `json:"icon" binding:"required"`
	IsDefault bool `json:"is_default"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Color string `json:"color"`
	Icon string `json:"icon"`
	IsDefault bool `json:"is_default"`
}

func (cs *CategoryService) CategoryCreate(user_id uuid.UUID, req CategoryCreateRequest) error {
	cat := models.Category{
		UserID: user_id,
		Name: req.Name,
		Type: req.Type,
		Color: req.Color,
		Icon: req.Icon,
		IsDefault: req.IsDefault,
	}
	if err := cs.category_repo.CreateCategory(&cat); err != nil {
		return err
	}
	return nil
}

func (cs *CategoryService) CategoryUpdate(user_id, cat_id uuid.UUID, req CategoryUpdateRequest) (*models.Category, error) {

}

func (cs *CategoryService) CategoryDelete(user_id, cat_id uuid.UUID) (*models.Category, error) {

}

func (cs *CategoryService) CategoryList(user_id uuid.UUID) ([]models.Category, error) {

}

func (cs *CategoryService) CategoryGet(user_id, cat_id uuid.UUID) (*models.Category, error) {

}
