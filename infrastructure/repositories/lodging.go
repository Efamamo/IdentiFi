package repositories

import (
	"fmt"

	"github.com/Efamamo/WonderBeam/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LodgingRepo struct {
	DB *gorm.DB
}

func (lr LodgingRepo) SaveLodging(lodging domain.Lodging) (*domain.Lodging, error) {
	var location domain.Location
	result := lr.DB.Where("id = ?", lodging.LocationID).First(&location)

	if result.Error != nil {
		return nil, result.Error
	}

	lodging.ID = uuid.New()

	result = lr.DB.Create(&lodging)

	if result.Error != nil {
		return nil, result.Error
	}
	return &lodging, nil
}

func (lr LodgingRepo) DeleteLodging(id string) error {
	result := lr.DB.Where("id = ?", id).Delete(&domain.Lodging{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with id: %s", id)
	}
	return nil
}
func (lr LodgingRepo) UpdateLodging(lid string, updateValues domain.LodgingUpdate) (*domain.Lodging, error) {
	var lodging domain.Lodging

	if updateValues.Name != "" {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("name", updateValues.Name)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.GoogleLink != "" {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("google_link", updateValues.GoogleLink)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.Description != "" {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("description", updateValues.Description)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.BudgetPerNight != 0 {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("budget_per_night", updateValues.BudgetPerNight)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.Category != "" {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("category", updateValues.Category)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.QualityRating != 0 {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("quality_rating", updateValues.QualityRating)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.UserRating != 0 {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("user_rating", updateValues.UserRating)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.Website != "" {
		result := lr.DB.Model(&domain.Lodging{}).Where("id = ?", lid).Update("website", updateValues.Website)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	result := lr.DB.Where("id = ?", lid).First(&lodging)
	if result.Error != nil {
		return nil, result.Error
	}

	return &lodging, nil
}
func (lr LodgingRepo) GetLodgings(locationId string) (*[]domain.Lodging, error) {
	var lodgings []domain.Lodging

	result := lr.DB.Where("location_id", locationId).Find(&lodgings)

	if result.Error != nil {
		return nil, result.Error
	}
	return &lodgings, nil
}

func (lr LodgingRepo) GetLodgingById(id string) (*domain.Lodging, error) {
	var lodging domain.Lodging

	result := lr.DB.Where("id = ?", id).First(&lodging)

	if result.Error != nil {
		return nil, result.Error
	}
	return &lodging, nil
}
