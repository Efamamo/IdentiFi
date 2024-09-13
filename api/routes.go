package api

import (
	"github.com/Efamamo/WonderBeam/api/controllers"
	"github.com/Efamamo/WonderBeam/infrastructure/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController, authController controllers.AuthController, lodgingController controllers.LodgingController, activityController controllers.ActivityController) {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	r.POST("/auth/signup", authController.Signup)
	r.POST("/auth/login", authController.Login)
	r.GET("/auth/verify", authController.VerifyEmail)

	r.GET("/locations", middlewares.VerifyTokenMiddleware(false), locationController.GetLocations)
	r.GET("/locations/:id", middlewares.VerifyTokenMiddleware(false), locationController.GetLocationById)
	r.POST("/locations", middlewares.VerifyTokenMiddleware(true), locationController.AddLocation)
	r.PATCH("/locations/:id", middlewares.VerifyTokenMiddleware(true), locationController.UpdateLocation)
	r.DELETE("/locations/:id", middlewares.VerifyTokenMiddleware(true), locationController.DeleteLocation)

	r.GET("/locations/:id/lodgings", middlewares.VerifyTokenMiddleware(false), lodgingController.GetLodgings)
	r.GET("/locations/:id/lodgings/:lid", middlewares.VerifyTokenMiddleware(false), lodgingController.GetLodgingById)
	r.POST("/locations/:id/lodgings", middlewares.VerifyTokenMiddleware(true), lodgingController.AddLodging)
	r.PATCH("/locations/:id/lodgings/:lid", middlewares.VerifyTokenMiddleware(true), lodgingController.UpdateLodging)
	r.DELETE("/locations/:id/lodgings/:lid", middlewares.VerifyTokenMiddleware(true), lodgingController.DeleteLodging)

	r.GET("/locations/:id/lodgings/:lid/activities", middlewares.VerifyTokenMiddleware(false), activityController.GetActivities)
	r.POST("/locations/:id/lodgings/:lid/activities", middlewares.VerifyTokenMiddleware(true), activityController.AddActivity)
	r.PATCH("/locations/:id/lodgings/:lid/activities/:aid", middlewares.VerifyTokenMiddleware(true), activityController.UpdateActivity)
	r.DELETE("/locations/:id/lodgings/:lid/activities/:aid", middlewares.VerifyTokenMiddleware(true), activityController.DeleteActivity)

	r.Run("localhost:8080")
}
