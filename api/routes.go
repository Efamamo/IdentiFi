package api

import (
	"github.com/Efamamo/IdentiFi/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController, authController controllers.AuthController, lodgingController controllers.LodgingController) {
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

	r.Run("localhost:5060")
}
