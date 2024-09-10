package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/location")
	r.POST("/location")
	r.PATCH("/location/:id")
	r.DELETE("/location/:id")

	r.Run()
}
