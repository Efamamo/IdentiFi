package main

import (
	"github.com/Efamamo/IdentiFi/api"
	"github.com/Efamamo/IdentiFi/api/controllers"
	usecase_interfaces "github.com/Efamamo/IdentiFi/api/interfaces"
	"github.com/Efamamo/IdentiFi/usecases"
)

func main() {
	var locationUsecase usecase_interfaces.ILocation = usecases.LocationUsecase{}
	locationController := controllers.LocationController{LocationUseCase: locationUsecase}
	var authUsecase usecase_interfaces.IAuthUsecase = usecases.AuthUsecase{}
	authController := controllers.AuthController{AuthUsecase: authUsecase}
	api.StartServer(locationController, authController)
}
