package api

import (
	"github.com/Efamamo/WonderBeam/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController, authController controllers.AuthController, lodgingController controllers.LodgingController, activityController controllers.ActivityController) {
	r := gin.Default()

	r.POST("/auth/signup", authController.Signup)
	r.POST("/auth/login", authController.Login)

	r.GET("/locations", locationController.GetLocations)
	r.POST("/locations", locationController.AddLocation)
	r.PATCH("/locations/:id", locationController.UpdateLocation)
	r.DELETE("/locations/:id", locationController.DeleteLocation)

	r.GET("/lodgings", lodgingController.GetLodgings)
	r.GET("/lodgings/:id", lodgingController.GetLodgingById)
	r.POST("/lodgings", lodgingController.AddLodging)
	r.PATCH("/lodgings/:id", lodgingController.UpdateLodging)
	r.DELETE("/lodgings/:id", lodgingController.DeleteLodging)

	r.GET("/activities", activityController.GetActivities)
	r.POST("/activities", activityController.AddActivity)
	r.PATCH("/activities/:id", activityController.UpdateActivity)
	r.DELETE("/activities/:id", activityController.DeleteActivity)

	r.Run("localhost:5060")
}
