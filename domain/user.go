package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Use UUID as the primary
	Username          string    `json:"username" binding:"required"`
	Email             string    `json:"email" binding:"required" gorm:"unique"`
	Password          string    `json:"password" binding:"required"`
	VerificationToken string    `json:"verification_token"`
	IsVerified        bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt         time.Time
	IsAdmin           bool `json:"is_admin"`
}
