package api

import (
	"github.com/Efamamo/IdentiFi/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController, authController controllers.AuthController) {
	r := gin.Default()

	r.POST("/auth/signup", authController.Signup)
	r.POST("/auth/login", authController.Login)

	r.GET("/locations", locationController.GetLocations)
	r.POST("/locations", locationController.AddLocation)
	r.PATCH("/locations/:id", locationController.UpdateLocation)
	r.DELETE("/locations/:id", locationController.DeleteLocation)

	r.Run("localhost:5050")
}
