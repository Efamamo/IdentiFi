package controllers

import (
	"net/http"

	"github.com/Efamamo/WonderBeam/api/controllers/dtos"
	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase usecase_interfaces.IAuthUsecase
}

func (ac AuthController) Signup(ctx *gin.Context) {
	user := domain.User{}

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ac.AuthUsecase.Signup(user)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Verify Your Email"})
}

func (ac AuthController) Login(ctx *gin.Context) {
	login := dtos.LoginDTO{}

	err := ctx.BindJSON(&login)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.AuthUsecase.Login(login.Username, login.Password)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	ctx.IndentedJSON(http.StatusCreated, token)

}
