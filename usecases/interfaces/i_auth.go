package repo_interfaces

import (
	"github.com/Efamamo/WonderBeam/domain"
)

type IAuthRepo interface {
	SignUp(domain.User) error
	FindUserByUsername(string) (*domain.User, error)
	VerifyEmail(string) error
}

type IJwtServices interface {
	GenerateAccessToken(domain.User) (string, error)
	GenerateRefreshToken(domain.User) (string, error)
}

type IPasswordServices interface {
	HashPassword(string) (string, error)
	MatchPassword(string, string) error
}
