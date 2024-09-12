package repo_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type ILodgingRepo interface {
	SaveLodging(domain.Lodging) (*domain.Lodging, error)
	DeleteLodging(string) error
	DeleteLocationLodgings(id string) error
	UpdateLodging(string, domain.LodgingUpdate) (*domain.Lodging, error)
	GetLodgings(string) (*[]domain.Lodging, error)
	GetLodgingById(string) (*domain.Lodging, error)
}
