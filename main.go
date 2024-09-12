package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Efamamo/WonderBeam/api"
	"github.com/Efamamo/WonderBeam/api/controllers"
	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/infrastructure"
	"github.com/Efamamo/WonderBeam/infrastructure/repositories"
	"github.com/Efamamo/WonderBeam/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infrastructure.ConnectToDB()
	infrastructure.Migrate()

	var authUsecase usecase_interfaces.IAuthUsecase = usecases.AuthUsecase{AuthRepo: repositories.AuthRepo{DB: infrastructure.DB}, JwtServices: infrastructure.Token{}, PasswordServices: infrastructure.Pass{}}
	authController := controllers.AuthController{AuthUsecase: authUsecase}

	var locationUsecase usecase_interfaces.ILocation = usecases.LocationUsecase{LocationRepo: repositories.LocationRepo{DB: infrastructure.DB}}
	locationController := controllers.LocationController{LocationUseCase: locationUsecase}

	var lodgingUsecase usecase_interfaces.ILodging = usecases.LodgingUsecase{LodgingRepo: repositories.LodgingRepo{DB: infrastructure.DB}}
	lodgingController := controllers.LodgingController{LodgingUsecase: lodgingUsecase}

	var activityUsecase usecase_interfaces.IActivity = usecases.ActivityUsecase{ActivityRepo: repositories.ActivityRepo{DB: infrastructure.DB}}
	activityController := controllers.ActivityController{ActivityUsecase: activityUsecase}

	api.StartServer(locationController, authController, lodgingController, activityController)
}
