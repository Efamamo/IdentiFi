package controllers

import (
	"net/http"

	usecase_interfaces "github.com/Efamamo/IdentiFi/api/interfaces"
	"github.com/Efamamo/IdentiFi/domain"
	"github.com/gin-gonic/gin"
)

type LodgingController struct {
	LodgingUsecase usecase_interfaces.ILodging
}

func (ldc LodgingController) GetLodgings(ctx *gin.Context) {
	location := ctx.Query("location")

	lodgings, err := ldc.LodgingUsecase.GetLodgings(location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, lodgings)

}

func (ldc LodgingController) GetLodgingById(ctx *gin.Context) {
	id := ctx.Param("id")

	lodging, err := ldc.LodgingUsecase.GetLodgingById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, lodging)

}

func (ldc LodgingController) AddLodging(ctx *gin.Context) {
	lodging := domain.Lodging{}

	err := ctx.BindJSON(&lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	new_lodging, err := ldc.LodgingUsecase.AddLodging(lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, new_lodging)
}

func (ldc LodgingController) UpdateLodging(ctx *gin.Context) {
	id := ctx.Param("id")

	updateLodging := domain.LodgingUpdate{}

	err := ctx.BindJSON(&updateLodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updatedLodging, err := ldc.LodgingUsecase.UpdateLodging(id, updateLodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedLodging)
}

func (ldc LodgingController) DeleteLodging(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ldc.LodgingUsecase.DeleteLodging(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}
