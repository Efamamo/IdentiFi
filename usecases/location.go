package usecases

import (
	"github.com/Efamamo/WonderBeam/domain"
	repo_interfaces "github.com/Efamamo/WonderBeam/usecases/interfaces"
)

type LocationUsecase struct {
	LocationRepo repo_interfaces.ILocationRepo
	LodgingRepo  repo_interfaces.ILodgingRepo
}

func (lu LocationUsecase) GetLocations() (*[]domain.Location, error) {
	locations, err := lu.LocationRepo.Get()

	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (lu LocationUsecase) AddLocation(l domain.Location) (*domain.Location, error) {
	location, err := lu.LocationRepo.Save(l)

	if err != nil {
		return nil, err
	}

	return location, nil
}

func (lu LocationUsecase) UpdateLocation(id string, l domain.LocationUpdate) (*domain.Location, error) {
	updatedLocation, err := lu.LocationRepo.Update(id, l)

	if err != nil {
		return nil, err
	}

	return updatedLocation, nil
}

func (lu LocationUsecase) DeleteLocation(id string) error {
	err := lu.LocationRepo.Delete(id)

	if err != nil {
		return err
	}

	err = lu.LodgingRepo.DeleteLocationLodgings(id)

	if err != nil {
		return err
	}

	return nil
}

func (lu LocationUsecase) GetLocationById(id string) (*domain.Location, error) {
	location, err := lu.LocationRepo.GetLocationById(id)

	if err != nil {
		return nil, err
	}

	return location, nil
}
