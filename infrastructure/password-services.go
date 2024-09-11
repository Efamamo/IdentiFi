package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type Pass struct{}

func (p Pass) HashPassword(password string) (string, error) {
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if e != nil {
		return "", e
	}
	return string(hashedPassword), nil
}

func (p Pass) MatchPassword(euPassword string, uPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(euPassword), []byte(uPassword))

	if err != nil {
		return err
	}

	return nil
}
