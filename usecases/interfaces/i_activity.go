package repo_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type IActivityRepo interface {
	SaveActivity(domain.Activity) (*domain.Activity, error)
	DeleteActivity(string) error
	UpdateActivity(string, domain.ActivityUpdate) (*domain.Activity, error)
	GetActivities(string) (*[]domain.Activity, error)
}
