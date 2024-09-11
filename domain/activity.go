package domain

import "time"

type Activity struct {
	Name     string    `json:"name" binding:"required"`
	Date     time.Time `json:"date" binding:"required"`
	ImageURL string    `json:"image_url" binding:"required"`
}

type ActivityUpdate struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	ImageURL string    `json:"image_url"`
}
