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
	ImageURL   string    `binding:"required" json:"image_url"`
	Lodgings   []Lodging `json:"lodgings" gorm:"constraint:OnDelete:CASCADE;"`
}

type LocationUpdate struct {
	Name       string `json:"name"`
	GoogleLink string `json:"google_link"`
	ImageURL   string `json:"image_url"`
}
