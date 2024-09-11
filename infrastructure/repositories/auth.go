package repositories

import (
	"errors"

	"github.com/Efamamo/WonderBeam/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepo struct{ DB *gorm.DB }

func (ar AuthRepo) SignUp(user domain.User) error {
	u, err := ar.FindUserByEmail(user.Email)

	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("email taken")
	}

	u, err = ar.FindUserByUsername(user.Username)

	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("username taken")
	}

	user.Id = uuid.New()

	result := ar.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar AuthRepo) FindUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := ar.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		} else {
			return nil, nil
		}
	}

	return &user, nil
}
func (ar AuthRepo) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := ar.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		} else {
			return nil, nil
		}
	}

	return &user, nil
}
