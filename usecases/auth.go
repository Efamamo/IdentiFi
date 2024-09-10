package usecases

import (
	"github.com/Efamamo/IdentiFi/domain"
	repo_interfaces "github.com/Efamamo/IdentiFi/usecases/interfaces"
)

type AuthUsecase struct {
	AuthRepo         repo_interfaces.IAuthRepo
	JwtServices      repo_interfaces.IJwtServices
	PasswordServices repo_interfaces.IPasswordServices
}

func (au AuthUsecase) Signup(user domain.User) error {
	hashedPassword, err := au.PasswordServices.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = au.AuthRepo.SignUp(user)
	if err != nil {
		return err
	}

	return nil

}

func (au AuthUsecase) Login(username string, password string) (*domain.Token, error) {
	user, err := au.AuthRepo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = au.PasswordServices.MatchPassword(password, user.Password)
	if err != nil {
		return nil, err
	}

	accessToken, err := au.JwtServices.GenerateAccessToken(*user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := au.JwtServices.GenerateRefreshToken(*user)
	if err != nil {
		return nil, err
	}

	token := domain.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &token, nil

}
