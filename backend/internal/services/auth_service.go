package services

import (
	"errors"

	"github.com/f18charles/piggy-bank/backend/internal/auth"
	"github.com/f18charles/piggy-bank/backend/internal/models"
	"github.com/f18charles/piggy-bank/backend/internal/repository"
	"github.com/f18charles/piggy-bank/backend/internal/utils"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

func (as *AuthService) RegisterUser(regreq RegisterRequest) (*models.User, string, error) {
	// check if email is taken
	_, err := as.userRepo.GetUserByEmail(regreq.Email)
	if err != nil {
		return nil, "", utils.ErrAlreadyExists
	}
	if !errors.Is(err, utils.ErrNotFound) {
		return nil, "", err
	}

	//hash password
	passHash, err := auth.HashPassword(regreq.Password)
	if err != nil {
		return nil, "", err
	}

	// create the user
	user := &models.User{
		Email:        regreq.Email,
		PasswordHash: passHash,
		FullName:     regreq.FullName,
		Currency:     "KES",
	}

	// save to db
	if err := as.userRepo.CreateUser(user); err != nil {
		return nil, "", err
	}

	// generate session token
	token, err := auth.GenarateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (as *AuthService) LoginUser(logreq LoginRequest) (*models.User, string, error) {
	// get user by email
	user, err := as.userRepo.GetUserByEmail(logreq.Email)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return nil, "", utils.ErrUnauthorized
		}
		return nil, "", err
	}

	// confirm password
	if !auth.CheckPass(logreq.Password, user.PasswordHash) {
		return nil, "", utils.ErrUnauthorized
	}

	token, err := auth.GenarateToken(user.ID)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (as *AuthService) GetAuthedUser(userID uuid.UUID) (*models.User, error) {
	return as.userRepo.GetUserByID(userID)
}
