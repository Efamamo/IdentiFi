package usecase_interfaces

import "github.com/Efamamo/IdentiFi/domain"

type ILocation interface {
	GetLocations() (*[]domain.Location, error)
	AddLocation(domain.Location) (*domain.Location, error)
	UpdateLocation(string, domain.Location) (*domain.Location, error)
	DeleteLocation(string) error
}
