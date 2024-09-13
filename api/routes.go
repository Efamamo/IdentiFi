package api

import (
	"github.com/Efamamo/WonderBeam/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController, authController controllers.AuthController, lodgingController controllers.LodgingController, activityController controllers.ActivityController) {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	r.POST("/auth/signup", authController.Signup)
	r.POST("/auth/login", authController.Login)
	r.GET("/auth/verify", authController.VerifyEmail)

	r.GET("/locations", locationController.GetLocations)
	r.GET("/locations/:id", locationController.GetLocationById)
	r.POST("/locations", locationController.AddLocation)
	r.PATCH("/locations/:id", locationController.UpdateLocation)
	r.DELETE("/locations/:id", locationController.DeleteLocation)

	r.GET("/locations/:id/lodgings", lodgingController.GetLodgings)
	r.GET("/locations/:id/lodgings/:lid", lodgingController.GetLodgingById)
	r.POST("/locations/:id/lodgings", lodgingController.AddLodging)
	r.PATCH("/locations/:id/lodgings/:lid", lodgingController.UpdateLodging)
	r.DELETE("/locations/:id/lodgings/:lid", lodgingController.DeleteLodging)

	r.GET("/locations/:id/lodgings/:lid/activities", activityController.GetActivities)
	r.POST("/locations/:id/lodgings/:lid/activities", activityController.AddActivity)
	r.PATCH("/locations/:id/lodgings/:lid/activities/:aid", activityController.UpdateActivity)
	r.DELETE("/locations/:id/lodgings/:lid/activities/:aid", activityController.DeleteActivity)

	r.Run("localhost:8080")
}
