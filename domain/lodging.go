package domain

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Lodging struct {
	gorm.Model
	ID             uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Use UUID as the primary
	Name           string         `binding:"required" json:"name"`
	GoogleLink     string         `binding:"required" json:"google_link"`
	Description    string         `binding:"required" json:"description"`
	BudgetPerNight float64        `binding:"required" json:"budget_per_night"`
	Category       string         `binding:"required" json:"category"`
	QualityRating  int            `binding:"required" json:"quality_rating"`
	UserRating     int            `binding:"required" json:"user_rating"`
	Emails         pq.StringArray `binding:"required" json:"emails" gorm:"type:text[]"`        // pq.StringArray
	PhoneNumbers   pq.StringArray `binding:"required" json:"phone_numbers" gorm:"type:text[]"` // pq.StringArray
	Website        string         `binding:"required" json:"website"`
	Amenities      pq.StringArray `binding:"required" json:"amenities" gorm:"type:text[]"` // pq.StringArray
	LocationID     uuid.UUID      `json:"location_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	Image          string         `json:"image"`
}

type LodgingUpdate struct {
	Name           string         `json:"name"`
	GoogleLink     string         `json:"google_link"`
	Description    string         `json:"description"`
	BudgetPerNight float64        `json:"budget_per_night"`
	Category       string         `json:"category"`
	QualityRating  int            `json:"quality_rating"`
	UserRating     int            `json:"user_rating"`
	Website        string         `json:"website"`
	Image          string         `json:"image"`
	Emails         pq.StringArray `binding:"required" json:"emails" gorm:"type:text[]"`        // pq.StringArray
	PhoneNumbers   pq.StringArray `binding:"required" json:"phone_numbers" gorm:"type:text[]"` // pq.StringArray
	Amenities      pq.StringArray `binding:"required" json:"amenities" gorm:"type:text[]"`     // pq.StringArray

}
