package repo_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type IActivityRepo interface {
	Save(domain.Activity) (*domain.Activity, error)
	Delete(string) error
	Update(string, domain.ActivityUpdate) (*domain.Activity, error)
	Get(string) (*[]domain.Activity, error)
}
