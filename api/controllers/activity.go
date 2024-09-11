package controllers

import (
	"net/http"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityUsecase usecase_interfaces.IActivity
}

func (ac ActivityController) AddActivity(ctx *gin.Context) {
	activity := domain.Activity{}

	err := ctx.BindJSON(&activity)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	loc, err := ac.ActivityUsecase.AddActivity(activity)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, *loc)
}

func (ac ActivityController) UpdateActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	act := domain.ActivityUpdate{}

	err := ctx.BindJSON(&act)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updatedActivity, err := ac.ActivityUsecase.UpdateActivity(id, act)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, *updatedActivity)
}

func (ac ActivityController) GetActivities(ctx *gin.Context) {

	lodging := ctx.Query("lodging")

	activities, err := ac.ActivityUsecase.GetActivities(lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.IndentedJSON(http.StatusAccepted, activities)
}

func (ac ActivityController) DeleteActivity(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ac.ActivityUsecase.DeleteActivity(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
	}

	ctx.Status(http.StatusNoContent)
}
