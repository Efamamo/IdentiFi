package repo_interfaces

import "github.com/Efamamo/IdentiFi/domain"

type ILocationRepo interface {
	Save(domain.Location) (*domain.Location, error)
	Delete(string) error
	Update(string, domain.Location) (*domain.Location, error)
	Get() (*[]domain.Location, error)
}
