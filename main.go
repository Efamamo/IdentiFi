package main

import (
	"github.com/Efamamo/IdentiFi/api"
	"github.com/Efamamo/IdentiFi/api/controllers"
)

func main() {
	locationController := controllers.LocationController{}
	api.StartServer(locationController)
}
