package controllers

import (
	"fmt"
	"net/http"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ActivityController struct {
	ActivityUsecase usecase_interfaces.IActivity
}

func (ac ActivityController) AddActivity(ctx *gin.Context) {
	activity := domain.Activity{}
	lodging := ctx.Param("lid")
	lodgingId, err := uuid.Parse(lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activity.LodgingId = lodgingId
	fmt.Println(activity)

	err = ctx.BindJSON(&activity)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loc, err := ac.ActivityUsecase.AddActivity(activity)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, *loc)
}

func (ac ActivityController) UpdateActivity(ctx *gin.Context) {
	id := ctx.Param("aid")
	act := domain.ActivityUpdate{}

	err := ctx.BindJSON(&act)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedActivity, err := ac.ActivityUsecase.UpdateActivity(id, act)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, *updatedActivity)
}

func (ac ActivityController) GetActivities(ctx *gin.Context) {

	lodging := ctx.Param("lid")

	activities, err := ac.ActivityUsecase.GetActivities(lodging)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusAccepted, activities)
}

func (ac ActivityController) DeleteActivity(ctx *gin.Context) {
	id := ctx.Param("aid")
	err := ac.ActivityUsecase.DeleteActivity(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	ctx.Status(http.StatusNoContent)
}
