package handlers

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/auth"
	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GoalHandler struct {
	goalService services.GoalService
}

func NewGoalHandler() *GoalHandler {
	return &GoalHandler{
		goalService: *services.NewGoalService(),
	}
}

// ListGoals returns all goals for the authenticated user.
func (gh *GoalHandler) ListGoals(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	goal, err := gh.goalService.GoalList(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "no goals found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, goal)
}

// CreateGoal creates a new savings goal for the authenticated user.
func (gh *GoalHandler) CreateGoal(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var goalreq services.GoalCreateRequest
	if err := c.ShouldBindJSON(&goalreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	goal, err := gh.goalService.GoalCreate(id, goalreq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create account")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, goal)
}

// GetGoal returns a single goal by id for the authenticated user.
func (gh *GoalHandler) GetGoal(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	param_id := c.Param("id")
	goal_id, err := uuid.Parse(param_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid goal id")
		return
	}

	goal, err := gh.goalService.GetGoal(id, goal_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "goal not found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, goal)
}

// UpdateGoal updates fields on a user's goal.
func (gh *GoalHandler) UpdateGoal(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	param_id := c.Param("id")
	goal_id, err := uuid.Parse(param_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid goal id")
		return
	}
	var goalreq services.GoalUpdateRequest
	if err := c.ShouldBindJSON(&goalreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	goal, err := gh.goalService.GoalUpdate(id, goal_id, goalreq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update account")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, goal)
}

// DeleteGoal deletes a user's goal.
func (gh *GoalHandler) DeleteGoal(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	param_id := c.Param("id")
	goal_id, err := uuid.Parse(param_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid goal id")
		return
	}
	err = gh.goalService.GoalDelete(id, goal_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete goal")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "goal deleted successfully"})
}
