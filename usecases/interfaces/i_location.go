package repo_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type ILocationRepo interface {
	Save(domain.Location) (*domain.Location, error)
	Delete(string) error
	Update(string, domain.LocationUpdate) (*domain.Location, error)
	Get() (*[]domain.Location, error)
	GetLocationById(string) (*domain.Location, error)
}
