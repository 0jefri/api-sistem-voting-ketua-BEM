package dto

import (
	"time"
)

type UserResponse struct {
	ID       string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Username string `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required,alphanum"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email,omitempty" binding:"required,email"`
	Role     string `gorm:"type:varchar(50);not null;default:'mahasiswa'" json:"role" binding:"required"`

	RegistrationDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"registrationDate" binding:"omitempty"`
	LastLogin        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastLogin" binding:"omitempty"`
}
