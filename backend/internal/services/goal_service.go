package services

import (
	"time"

	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
)

type GoalCreateRequest struct {
	Name          string     `json:"name" binding:"required"`
	TargetAmount  float64    `json:"target_amount" binding:"required"`
	CurrentAmount float64    `json:"current_amount"`
	Deadline      *time.Time `json:"deadline"`
}

type GoalService struct {
	goalRepo *repository.GoalRepo
}

// NewGoalService creates a GoalService with an initialized repository.
func NewGoalService() *GoalService {
	return &GoalService{
		goalRepo: repository.NewGoalRepo(),
	}
}

type GoalUpdateRequest struct {
	Name          string     `json:"name"`
	TargetAmount  float64    `json:"target_amount"`
	CurrentAmount float64    `json:"current_amount"`
	Deadline      *time.Time `json:"deadline"`
}

// GoalCreate creates a new savings goal for the user and persists it via the repository.
func (gs *GoalService) GoalCreate(user_id uuid.UUID, req GoalCreateRequest) (*models.Goal, error) {
	goal := &models.Goal{
		UserID:        user_id,
		Name:          req.Name,
		TargetAmount:  req.TargetAmount,
		CurrentAmount: req.CurrentAmount,
		Deadline:      req.Deadline,
	}
	if err := gs.goalRepo.CreateGoal(goal); err != nil {
		return nil, err
	}
	return goal, nil
}

// GetGoal fetches a goal by ID and ensures the requesting user is the owner.
func (gs *GoalService) GetGoal(user_id, goal_id uuid.UUID) (*models.Goal, error) {
	goal, err := gs.goalRepo.GetGoalByID(goal_id)
	if err != nil {
		return nil, err
	}
	if goal.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	return goal, nil
}

// GoalUpdate updates an existing goal's fields that were provided in the request.
func (gs *GoalService) GoalUpdate(user_id, goal_id uuid.UUID, req GoalUpdateRequest) (*models.Goal, error) {
	goal, err := gs.goalRepo.GetGoalByID(goal_id)
	if err != nil {
		return nil, err
	}
	if goal.UserID != user_id {
		return nil, utils.ErrForbidden
	}
	if req.Name != "" {
		goal.Name = req.Name
	}
	if req.TargetAmount != 0 {
		goal.TargetAmount = req.TargetAmount
	}
	if req.CurrentAmount != 0 {
		goal.CurrentAmount = req.CurrentAmount
	}
	if req.Deadline != nil {
		goal.Deadline = req.Deadline
	}
	if err := gs.goalRepo.UpdateGoal(goal); err != nil {
		return nil, err
	}
	return goal, nil
}

// GoalList returns all goals for the specified user.
func (gs *GoalService) GoalList(user_id uuid.UUID) ([]models.Goal, error) {
	goals, err := gs.goalRepo.ListGoalsByUser(user_id)
	if err != nil {
		return nil, err
	}
	return goals, nil
}

// GoalDelete deletes a goal after verifying ownership.
func (gs *GoalService) GoalDelete(user_id, goal_id uuid.UUID) error {
	goal, err := gs.goalRepo.GetGoalByID(goal_id)
	if err != nil {
		return err
	}
	if goal.UserID != user_id {
		return utils.ErrForbidden
	}
	return nil
}
