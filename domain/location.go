package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `binding:"required" json:"name"`
	GoogleLink string    `binding:"required" json:"google_link"`
	Image      string    `binding:"required" json:"image"`
}

type LocationUpdate struct {
	Name       string `json:"name"`
	GoogleLink string `json:"google_link"`
	Image      string `json:"image"`
}
