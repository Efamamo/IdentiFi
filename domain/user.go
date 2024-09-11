package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                uuid.UUID `json:"id" gorm:"primaryKey"`
	Username          string    `json:"username" binding:"required"`
	Email             string    `json:"email" binding:"required" gorm:"unique"`
	Password          string    `json:"password" binding:"required"`
	VerificationToken string    `json:"verification_token"`
	TimeOut           time.Time `json:"time_out"`
}
