package infrastructure

import (
	"os"
	"time"

	"github.com/Efamamo/WonderBeam/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct{}

func (tok Token) GenerateAccessToken(user domain.User) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"isAdmin":  user.IsAdmin,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))

	if e != nil {
		return "", e
	}

	return jwtToken, nil
}

func (tok Token) GenerateRefreshToken(user domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	jwtToken, e := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))
	if e != nil {
		return "", e
	}

	return jwtToken, nil
}
