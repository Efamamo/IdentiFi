package domain

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Id         string `json:"id" gorm:"primaryKey"`
	Name       string `binding:"required" json:"name"`
	GoogleLink string `binding:"required" json:"google_link"`
	ImageURL   string `binding:"required" json:"image_url"`
}

type LocationUpdate struct {
	Name       string `json:"name"`
	GoogleLink string `json:"google_link"`
	ImageURL   string `json:"image_url"`
}
