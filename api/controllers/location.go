package controllers

import (
	"net/http"

	usecase_interfaces "github.com/Efamamo/IdentiFi/api/interfaces"
	"github.com/Efamamo/IdentiFi/domain"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	LocationUseCase usecase_interfaces.ILocation
}

func (lc LocationController) AddLocation(ctx *gin.Context) {
	location := domain.Location{}

	err := ctx.BindJSON(&location)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	loc, err := lc.LocationUseCase.AddLocation(location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, *loc)
}

func (lc LocationController) UpdateLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	location := domain.LocationUpdate{}

	err := ctx.BindJSON(&location)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	loc, err := lc.LocationUseCase.UpdateLocation(id, location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, *loc)
}

func (lc LocationController) GetLocations(ctx *gin.Context) {

	locations, err := lc.LocationUseCase.GetLocations()

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.IndentedJSON(http.StatusAccepted, locations)
}

func (lc LocationController) DeleteLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	err := lc.LocationUseCase.DeleteLocation(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
	}

	ctx.Status(http.StatusNoContent)
}
