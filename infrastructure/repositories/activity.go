package repositories

import (
	"fmt"

	"github.com/Efamamo/WonderBeam/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivityRepo struct {
	DB *gorm.DB
}

func (ar ActivityRepo) SaveActivity(activity domain.Activity) (*domain.Activity, error) {
	var lodging domain.Lodging

	result := ar.DB.Where("id = ?", activity.LodgingId).First(&lodging)

	if result.Error != nil {
		return nil, result.Error
	}

	activity.ID = uuid.New()
	result = ar.DB.Create(&activity)

	if result.Error != nil {
		return nil, result.Error
	}
	return &activity, nil
}

func (ar ActivityRepo) DeleteActivity(id string) error {
	result := ar.DB.Where("id = ?", id).Delete(&domain.Activity{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with id: %s", id)
	}
	return nil
}
func (ar ActivityRepo) UpdateActivity(aid string, updateValues domain.ActivityUpdate) (*domain.Activity, error) {
	var activity domain.Activity
	if updateValues.Name != "" {
		result := ar.DB.Model(&domain.Activity{}).Where("id = ?", aid).Update("name", updateValues.Name)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	if updateValues.ImageURL != "" {
		result := ar.DB.Model(&domain.Activity{}).Where("id = ?", aid).Update("image_url", updateValues.ImageURL)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	result := ar.DB.Where("id = ?", aid).First(&activity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activity, nil
}
func (ar ActivityRepo) GetActivities(lid string) (*[]domain.Activity, error) {
	var activities []domain.Activity
	result := ar.DB.Where("lodging_id = ?", lid).Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}
	return &activities, nil
}
