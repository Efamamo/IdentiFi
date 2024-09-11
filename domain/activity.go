package domain

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Id       string    `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
	ImageURL string    `json:"image_url" binding:"required"`
}

type ActivityUpdate struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	ImageURL string    `json:"image_url"`
}
