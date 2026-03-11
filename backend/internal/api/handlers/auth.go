package handlers

import (
	"errors"
	"net/http"

	"github.com/f18charles/piggy-bank/backend/internal/services"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: *services.NewAuthService(db),
	}
}

type AuthResponse struct {
	Token string `json:"token"`
	User  any    `json:"user"`
}

// Register handles POST /auth/register: validates request, registers a new
// user via AuthService and returns the created user and JWT.
func (ah *AuthHandler) Register(c *gin.Context) {
	var regreq services.RegisterRequest
	if err := c.ShouldBindJSON(&regreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := ah.authService.RegisterUser(regreq)
	if err != nil {
		if errors.Is(err, utils.ErrAlreadyExists) {
			utils.ErrorResponse(c, http.StatusConflict, "email used in other account")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create account")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, AuthResponse{
		Token: token,
		User:  user,
	})
}

// Login handles POST /auth/login: validates credentials and returns a JWT on
// success.
func (ah *AuthHandler) Login(c *gin.Context) {
	var logreq services.LoginRequest
	if err := c.ShouldBindJSON(&logreq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := ah.authService.LoginUser(logreq)
	if err != nil {
		if errors.Is(err, utils.ErrUnauthorized) {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to login")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}

// Logout returns a success response for client-side logout flows.
func (ah *AuthHandler) Logout(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "Logged Out Successfully"})
}

// Profile returns the authenticated user's profile (GET /auth/profile).
func (ah *AuthHandler) Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, ok := userID.(uuid.UUID)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "invalid user id in context")
		return
	}

	user, err := ah.authService.GetAuthedUser(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, user)
}
