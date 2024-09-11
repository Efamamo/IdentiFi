package repositories

import (
	"fmt"

	"github.com/Efamamo/WonderBeam/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepo struct {
	DB *gorm.DB
}

func (lr LocationRepo) Save(location domain.Location) (*domain.Location, error) {
	location.Id = uuid.New()

	result := lr.DB.Create(&location)

	if result.Error != nil {
		return nil, result.Error
	}

	return &location, nil
}
func (lr LocationRepo) Get() (*[]domain.Location, error) {
	var locations []domain.Location

	result := lr.DB.Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}
	return &locations, nil
}
func (lr LocationRepo) Delete(id string) error {
	result := lr.DB.Where("id = ?", id).Delete(&domain.Location{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with id: %s", id)
	}
	return nil

}
func (lr LocationRepo) Update(id string, updateData domain.LocationUpdate) (*domain.Location, error) {
	var location domain.Location
	if updateData.GoogleLink != "" {
		result := lr.DB.Model(&domain.Location{}).Where("id = ?", id).Update("google_link", updateData.GoogleLink)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateData.ImageURL != "" {
		result := lr.DB.Model(&domain.Location{}).Where("id = ?", id).Update("image_url", updateData.ImageURL)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateData.Name != "" {
		result := lr.DB.Model(&domain.Location{}).Where("id = ?", id).Update("name", updateData.Name)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	result := lr.DB.Where("id = ?", id).First(&location)
	if result.Error != nil {
		return nil, result.Error
	}

	return &location, nil
}
