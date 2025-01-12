package service

import (
	"time"

	"github.com/api-voting/internal/app/repository"
	"github.com/api-voting/internal/model"
	"github.com/api-voting/internal/model/dto"
	"github.com/api-voting/utils/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*dto.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterNewUser(payload *model.User) (*dto.UserResponse, error) {

	users, err := s.repo.List()

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	for _, user := range users {
		if user.Username == payload.Username {
			return nil, exception.ErrUsernameAlreadyExist
		}
		if user.Email == payload.Email {
			return nil, exception.ErrEmailAlreadyExist
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	password := string(bytes)

	payload.Password = password

	user, err := s.repo.Create(payload)

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Role:             user.Role,
		RegistrationDate: user.RegistrationDate,
		LastLogin:        time.Now(),
	}

	return &userResponse, err
}
