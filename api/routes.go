package api

import (
	"github.com/Efamamo/IdentiFi/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController) {
	r := gin.Default()

	r.GET("/locations", locationController.GetLocations)
	r.POST("/locations", locationController.AddLocation)
	r.PATCH("/locations/:id", locationController.UpdateLocation)
	r.DELETE("/locations/:id", locationController.DeleteLocation)

	r.Run("localhost:5050")
}
