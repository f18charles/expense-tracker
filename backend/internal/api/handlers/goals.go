package handlers

import (
	"net/http"

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

func (gh *GoalHandler) ListGoals(c *gin.Context)  {
	id := utils.ConfirmAuthedUser(c)
	if id == uuid.Nil {
		return
	}
	goal, err := gh.goalService.GoalList(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "no goals found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, goal)
}

func (gh *GoalHandler) CreateGoal(c *gin.Context) {
	id := utils.ConfirmAuthedUser(c)
	if id == uuid.Nil {
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

func (gh *GoalHandler) GetGoal(c *gin.Context) {
	id := utils.ConfirmAuthedUser(c)
	if id == uuid.Nil {
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

func UpdateGoal(c *gin.Context) {

}

func DeleteGoal(c *gin.Context) {

}

