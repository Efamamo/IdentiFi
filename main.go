package main

import (
	_ "github.com/lib/pq"

	"github.com/Efamamo/WonderBeam/api"
	"github.com/Efamamo/WonderBeam/api/controllers"
	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/Efamamo/WonderBeam/infrastructure"
	"github.com/Efamamo/WonderBeam/infrastructure/repositories"
	"github.com/Efamamo/WonderBeam/usecases"
)

func main() {

	infrastructure.ConnectToDB()
	infrastructure.DB.AutoMigrate(&domain.User{})
	infrastructure.DB.AutoMigrate(&domain.Location{})
	infrastructure.DB.AutoMigrate(&domain.Lodging{})
	infrastructure.DB.AutoMigrate(&domain.Activity{})

	var authUsecase usecase_interfaces.IAuthUsecase = usecases.AuthUsecase{AuthRepo: repositories.AuthRepo{DB: infrastructure.DB}, JwtServices: infrastructure.Token{}, PasswordServices: infrastructure.Pass{}}
	authController := controllers.AuthController{AuthUsecase: authUsecase}

	var locationUsecase usecase_interfaces.ILocation = usecases.LocationUsecase{}
	locationController := controllers.LocationController{LocationUseCase: locationUsecase}

	var lodgingUsecase usecase_interfaces.ILodging = usecases.LodgingUsecase{}
	lodgingController := controllers.LodgingController{LodgingUsecase: lodgingUsecase}

	var activityUsecase usecase_interfaces.IActivity = usecases.ActivityUsecase{}
	activityController := controllers.ActivityController{ActivityUsecase: activityUsecase}

	api.StartServer(locationController, authController, lodgingController, activityController)
}
