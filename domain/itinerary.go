package domain

import "github.com/google/uuid"

type Itinerary struct {
	ID           uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	LocationName string     `json:"location_name"`
	LodgingName  string     `json:"lodging_name"`
	Activities   []Activity `json:"activities"`
}
