package usecases

import (
	"errors"

	"github.com/Efamamo/WonderBeam/domain"
	repo_interfaces "github.com/Efamamo/WonderBeam/usecases/interfaces"
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
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = au.PasswordServices.MatchPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
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
