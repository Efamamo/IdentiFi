package usecases

import (
	"github.com/Efamamo/IdentiFi/domain"
	repo_interfaces "github.com/Efamamo/IdentiFi/usecases/interfaces"
)

type LodgingUsecase struct {
	LodgingRepo repo_interfaces.ILodgingRepo
}

func (ldu LodgingUsecase) GetLodgings(location string) (*[]domain.Lodging, error) {
	lodgings, err := ldu.LodgingRepo.Get(location)

	if err != nil {
		return nil, err
	}

	return lodgings, nil
}

func (ldu LodgingUsecase) GetLodgingById(id string) (*domain.Lodging, error) {
	lodging, err := ldu.LodgingRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	return lodging, nil
}

func (ldu LodgingUsecase) AddLodging(lodging domain.Lodging) (*domain.Lodging, error) {
	l, err := ldu.LodgingRepo.Save(lodging)

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (ldu LodgingUsecase) UpdateLodging(id string, lodging domain.LodgingUpdate) (*domain.Lodging, error) {
	l, err := ldu.LodgingRepo.Update(id, lodging)

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (ldu LodgingUsecase) DeleteLodging(id string) error {
	err := ldu.LodgingRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
