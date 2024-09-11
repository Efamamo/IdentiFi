package repo_interfaces

import (
	"github.com/Efamamo/WonderBeam/api/controllers/dtos"
	"github.com/Efamamo/WonderBeam/domain"
)

type IAuthRepo interface {
	SignUp(domain.User) error
	LogIn(dtos.LoginDTO) (*domain.User, error)
	FindUserByUsername(string) (*domain.User, error)
}

type IJwtServices interface {
	GenerateAccessToken(domain.User) (string, error)
	GenerateRefreshToken(domain.User) (string, error)
}

type IPasswordServices interface {
	HashPassword(string) (string, error)
	MatchPassword(string, string) error
}
