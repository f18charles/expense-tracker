package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoalRepo struct {
	db *gorm.DB
}

func NewGoalRepo(db *gorm.DB) *GoalRepo {
	return &GoalRepo{
		db: db,
	}
}

func (gr *GoalRepo) CreateGoal(goal *models.Goal) error {
	return gr.db.Create(goal).Error
}

func (gr *GoalRepo) GetGoalByID(goalID uuid.UUID) (*models.Goal, error) {
	var goal models.Goal
	result := gr.db.Where("id = ?", goalID).First(&goal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &goal, nil
}

func (gr *GoalRepo) UpdateGoal(goal *models.Goal) error {
	return gr.db.Save(goal).Error
}

func (gr *GoalRepo) ListGoalsByUser(user_id uuid.UUID) ([]models.Goal, error) {
	goals := []models.Goal{}
	result := gr.db.Where("user_id = ?", user_id).Find(&goals)
	if result.Error != nil {
		return nil, result.Error
	}
	return goals, nil
}

func (gr *GoalRepo) DeleteGoal(id uuid.UUID) error {
	result := gr.db.Delete(&models.Goal{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
