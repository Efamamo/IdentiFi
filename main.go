package main

import (
	"github.com/Efamamo/WonderBeam/api"
	"github.com/Efamamo/WonderBeam/api/controllers"
	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/usecases"
)

func main() {

	var authUsecase usecase_interfaces.IAuthUsecase = usecases.AuthUsecase{}
	authController := controllers.AuthController{AuthUsecase: authUsecase}

	var locationUsecase usecase_interfaces.ILocation = usecases.LocationUsecase{}
	locationController := controllers.LocationController{LocationUseCase: locationUsecase}

	var lodgingUsecase usecase_interfaces.ILodging = usecases.LodgingUsecase{}
	lodgingController := controllers.LodgingController{LodgingUsecase: lodgingUsecase}

	var activityUsecase usecase_interfaces.IActivity = usecases.ActivityUsecase{}
	activityController := controllers.ActivityController{ActivityUsecase: activityUsecase}

	api.StartServer(locationController, authController, lodgingController, activityController)
}
