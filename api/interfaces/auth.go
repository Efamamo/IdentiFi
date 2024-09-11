package usecase_interfaces

import "github.com/Efamamo/WonderBeam/domain"

type IAuthUsecase interface {
	Signup(domain.User) error
	Login(string, string) (*domain.Token, error)
}
