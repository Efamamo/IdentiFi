package repositories

import (
	"errors"
	"fmt"
	"os"

	"github.com/Efamamo/WonderBeam/domain"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
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

	user.ID = uuid.New()
	user.VerificationToken = uuid.New().String()

	result := ar.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	ar.SendVerificationEmail(user)

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

func (ar AuthRepo) SendVerificationEmail(user domain.User) {

	from := "ephremmamo555@gmail.com"
	to := user.Email
	subject := "Verify your email"
	body := fmt.Sprintf(`
    <html>
    <body>
        <p>Hello %s,</p>
        <p>Please click the following link to verify your account:</p>
        <a href="http://localhost:8080/auth/verify?token=%s">Verify Your Account</a>
    </body>
    </html>`, user.Username, user.VerificationToken)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, os.Getenv("MAILTRAP_USERNAME"), os.Getenv("MAILTRAP_PASSWORD"))

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println("Could not send email:", err)
		return
	}

	fmt.Println("Verification email sent successfully")
}

func (ar AuthRepo) VerifyEmail(token string) error {

	var user domain.User
	if result := ar.DB.Where("verification_token = ?", token).First(&user); result.Error != nil {
		return result.Error
	}

	user.IsVerified = true
	user.VerificationToken = ""
	ar.DB.Save(&user)

	return nil
}
