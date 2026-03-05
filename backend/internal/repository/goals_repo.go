package repository

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/database"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoalRepo struct {}

func NewGoalRepo() *GoalRepo {
	return &GoalRepo{}
}

func (gr *GoalRepo) CreateGoal(goal *models.Goal) error {
	return database.DB.Create(goal).Error
}

func (gr *GoalRepo) GetGoalByID(goalID uuid.UUID) (*models.Goal, error) {
	var goal models.Goal
	result := database.DB.Where("id = ?", goalID).First(&goal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, result.Error
	}
	return &goal, nil
}

func (gr *GoalRepo) UpdateGoal(goal *models.Goal) error {
	return database.DB.Save(goal).Error
}

func ListGoalsByUser(user_id uuid.UUID) ([]models.Goal, error) {
	goals := []models.Goal{}
	result := database.DB.Where("user_id = ?", user_id).Find(&goals)
	if result.Error != nil {
		return nil, result.Error
	}
	return goals, nil
}

func (gr *GoalRepo) DeleteGoal(id uuid.UUID) error {
	result := database.DB.Delete(&models.Goal{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}