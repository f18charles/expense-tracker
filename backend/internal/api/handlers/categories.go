package handlers

import (
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/auth"
	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categoryService: services.NewCategoryService(),
	}
}

// ListCategories returns available categories (system and user-defined).
func (ch *CategoryHandler) ListCategories(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	categories, err := ch.categoryService.CategoryList(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "no categories found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, categories)
}

// GetCategory returns a single goal by id for the authenticated user.
func (ch *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cat_param := c.Param("id")
	cat_id, err := uuid.Parse(cat_param)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid category id")
		return
	}
	category, err := ch.categoryService.CategoryGet(id, cat_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "category not found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, category)
}

// CreateCategory creates a new category for the authenticated user.
func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var cat_req services.CategoryCreateRequest
	if err := c.ShouldBindJSON(&cat_req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	category, err := ch.categoryService.CategoryCreate(id, cat_req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create category")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, category)
}

// UpdateCategory updates a user's category.
func (ch *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	param_id := c.Param("id")
	cat_id, err := uuid.Parse(param_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid goal id")
		return
	}
	var cat_up_req services.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&cat_up_req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	category, err := ch.categoryService.CategoryUpdate(id, cat_id, cat_up_req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update category")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, category)
}

// DeleteCategory deletes a user's category.
func (ch *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := auth.ConfirmAuthedUser(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	param_id := c.Param("id")
	cat_id, err := uuid.Parse(param_id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid goal id")
		return
	}
	if err = ch.categoryService.CategoryDelete(id, cat_id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete category")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
