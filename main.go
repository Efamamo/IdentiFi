package main

import (
	"github.com/Efamamo/IdentiFi/api"
	"github.com/Efamamo/IdentiFi/api/controllers"
	usecase_interfaces "github.com/Efamamo/IdentiFi/api/interfaces"
	"github.com/Efamamo/IdentiFi/usecases"
)

func main() {
	var authUsecase usecase_interfaces.IAuthUsecase = usecases.AuthUsecase{}
	authController := controllers.AuthController{AuthUsecase: authUsecase}
	var locationUsecase usecase_interfaces.ILocation = usecases.LocationUsecase{}
	locationController := controllers.LocationController{LocationUseCase: locationUsecase}
	var lodgingUsecase usecase_interfaces.ILodging = usecases.LodgingUsecase{}
	lodgingController := controllers.LodgingController{LodgingUsecase: lodgingUsecase}

	api.StartServer(locationController, authController, lodgingController)
}
