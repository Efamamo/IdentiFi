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

	lodging.Id = uuid.New()

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
func (lr LodgingRepo) UpdateLodging(string, domain.LodgingUpdate) (*domain.Lodging, error) {
	return nil, nil
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
