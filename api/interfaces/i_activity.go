package usecase_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type IActivity interface {
	GetActivities(string) (*[]domain.Activity, error)
	AddActivity(domain.Activity) (*domain.Activity, error)
	UpdateActivity(string, domain.ActivityUpdate) (*domain.Activity, error)
	DeleteActivity(string) error
}
