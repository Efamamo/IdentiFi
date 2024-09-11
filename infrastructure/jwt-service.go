package infrastructure

import (
	"fmt"
	"os"
	"time"

	"github.com/Efamamo/WonderBeam/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct{}

func (tok Token) ValidateToken(t string) (*jwt.Token, error) {
	token, e := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JwtSecret")), nil
	})

	if e != nil || !token.Valid {
		return nil, e
	}
	return token, nil
}

func (tok Token) ValidateAdmin(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	role, ok := claims["isAdmin"].(bool)
	if !ok || !role {

		return false
	}
	return true
}

func (tok Token) GenerateAccessToken(user domain.User) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime,
	})

	jwtToken, e := token.SignedString([]byte("JwtSecret"))

	if e != nil {
		return "", e
	}

	return jwtToken, nil
}

func (tok Token) GenerateRefreshToken(user domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	jwtToken, e := token.SignedString([]byte("JwtSecret"))
	if e != nil {
		return "", e
	}

	return jwtToken, nil
}
