package usecases

import (
	"github.com/Efamamo/WonderBeam/domain"
	repo_interfaces "github.com/Efamamo/WonderBeam/usecases/interfaces"
)

type ActivityUsecase struct {
	ActivityRepo repo_interfaces.IActivityRepo
}

func (au ActivityUsecase) GetActivities(lodgingId string) (*[]domain.Activity, error) {
	activities, err := au.ActivityRepo.GetActivities(lodgingId)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (au ActivityUsecase) AddActivity(a domain.Activity) (*domain.Activity, error) {

	activity, err := au.ActivityRepo.SaveActivity(a)

	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (au ActivityUsecase) UpdateActivity(id string, a domain.ActivityUpdate) (*domain.Activity, error) {
	updatedActivity, err := au.ActivityRepo.UpdateActivity(id, a)

	if err != nil {
		return nil, err
	}

	return updatedActivity, nil
}

func (au ActivityUsecase) DeleteActivity(id string) error {
	err := au.ActivityRepo.DeleteActivity(id)

	if err != nil {
		return err
	}
	return nil
}
