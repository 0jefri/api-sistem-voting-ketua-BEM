package repository

import (
	"fmt"

	"github.com/api-voting/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[model.User]
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(payload *model.User) (*model.User, error) {
	user := model.User{
		ID:       payload.ID,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     payload.Role,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err

	}
	fmt.Println(user)

	return &user, nil
}

func (r *userRepository) List() ([]*model.User, error) {
	users := []*model.User{}

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
