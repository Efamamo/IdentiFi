package usecase_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type ILocation interface {
	GetLocations() (*[]domain.Location, error)
	AddLocation(domain.Location) (*domain.Location, error)
	UpdateLocation(string, domain.LocationUpdate) (*domain.Location, error)
	DeleteLocation(string) error
}
