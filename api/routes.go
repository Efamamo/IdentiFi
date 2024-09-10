package api

import (
	"github.com/Efamamo/IdentiFi/api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(locationController controllers.LocationController) {
	r := gin.Default()

	r.GET("/location", locationController.GetLocations)
	r.POST("/location", locationController.AddLocation)
	r.PATCH("/location/:id", locationController.UpdateLocation)
	r.DELETE("/location/:id", locationController.DeleteLocation)

	r.Run()
}
