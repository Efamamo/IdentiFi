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
	StartTime  time.Time `json:"start_time" binding:"required"` // Start of the activity
	EndTime    time.Time `json:"end_time" binding:"required"`   // End of the activity
	Image      string    `json:"image" binding:"required"`
	LodgingId  uuid.UUID `json:"lodging_id"`
	LocationId uuid.UUID `json:"location_id"`
}

type ActivityUpdate struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time" binding:"required"` // Start of the activity
	EndTime   time.Time `json:"end_time" binding:"required"`
	Image     string    `json:"image"`
}
