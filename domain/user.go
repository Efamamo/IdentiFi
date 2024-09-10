package domain

import "time"

type User struct {
	Id                string    `json:"id"`
	Username          string    `json:"username" binding:"required"`
	Email             string    `json:"email" binding:"required"`
	Password          string    `json:"password" binding:"required"`
	VerificationToken string    `json:"verification_token"`
	TimeOut           time.Time `json:"time_out"`
}
