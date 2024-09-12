package controllers

import (
	"net/http"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LodgingController struct {
	LodgingUsecase  usecase_interfaces.ILodging
}

func (ldc LodgingController) GetLodgings(ctx *gin.Context) {
	location := ctx.Param("id")

	lodgings, err := ldc.LodgingUsecase.GetLodgings(location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, lodgings)

}

func (ldc LodgingController) GetLodgingById(ctx *gin.Context) {
	id := ctx.Param("lid")

	lodging, err := ldc.LodgingUsecase.GetLodgingById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, lodging)

}

func (ldc LodgingController) AddLodging(ctx *gin.Context) {
	location := ctx.Param("id")
	lodging := domain.Lodging{}

	locationId, err := uuid.Parse(location)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lodging.LocationID = locationId

	err = ctx.BindJSON(&lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_lodging, err := ldc.LodgingUsecase.AddLodging(lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, new_lodging)
}

func (ldc LodgingController) UpdateLodging(ctx *gin.Context) {
	id := ctx.Param("lid")

	updateLodging := domain.LodgingUpdate{}

	err := ctx.BindJSON(&updateLodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLodging, err := ldc.LodgingUsecase.UpdateLodging(id, updateLodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedLodging)
}

func (ldc LodgingController) DeleteLodging(ctx *gin.Context) {
	id := ctx.Param("lid")

	err := ldc.LodgingUsecase.DeleteLodging(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
