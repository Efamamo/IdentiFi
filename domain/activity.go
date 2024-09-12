package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Use UUID as the primary
	Name       string    `json:"name" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	ImageURL   string    `json:"image_url" binding:"required"`
	LodgingId  uuid.UUID `json:"lodging_id"`
	LocationId uuid.UUID `json:"location_id"`
}

type ActivityUpdate struct {
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	ImageURL   string    `json:"image_url"`
	LodgingId  uuid.UUID `json:"lodging_id"`
	LocationId uuid.UUID `json:"location_id"`
}
