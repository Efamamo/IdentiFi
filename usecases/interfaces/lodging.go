package repo_interfaces

import "github.com/Efamamo/IdentiFi/domain"

type ILodgingRepo interface {
	Save(domain.Lodging) (*domain.Lodging, error)
	Delete(string) error
	Update(string, domain.LodgingUpdate) (*domain.Lodging, error)
	Get(string) (*[]domain.Lodging, error)
	GetById(string) (*domain.Lodging, error)
}
