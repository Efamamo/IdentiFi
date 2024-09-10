package usecase_interfaces

import "github.com/Efamamo/IdentiFi/domain"

type ILodging interface {
	GetLodgings(string) (*[]domain.Lodging, error)
	GetLodgingById(string) (*domain.Lodging, error)
	AddLodging(domain.Lodging) (*domain.Lodging, error)
	UpdateLodging(string, domain.LodgingUpdate) (*domain.Lodging, error)
	DeleteLodging(string) error
}
