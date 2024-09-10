package usecase_interfaces

import (
	"github.com/Efamamo/IdentiFi/api/controllers/dtos"
	"github.com/Efamamo/IdentiFi/domain"
)

type IAuthUsecase interface {
	Signup(domain.User) error
	Login(dtos.LoginDTO) (*domain.Token, error)
}
