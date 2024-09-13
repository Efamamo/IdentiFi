package controllers

import (
	"fmt"
	"net/http"
	"time"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ActivityController struct {
	ActivityUsecase usecase_interfaces.IActivity
}

func (ac ActivityController) AddActivity(ctx *gin.Context) {

	name := ctx.PostForm("name")
	if name == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	startTimeStr := ctx.PostForm("start_time")
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTimeStr := ctx.PostForm("end_time")
	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid end_time format"})
		return
	}

	lodging := ctx.Param("lid")
	lodgingId, err := uuid.Parse(lodging)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	locationId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	imagePath := fmt.Sprintf("./uploads/%s", file.Filename)

	err = ctx.SaveUploadedFile(file, imagePath)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid image format"})
		return
	}

	activity := domain.Activity{
		Name:       name,
		StartTime:  startTime,
		EndTime:    endTime,
		LodgingId:  lodgingId,
		LocationId: locationId,
		Image:      imagePath,
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
	name := ctx.PostForm("name")
	startTimeStr := ctx.PostForm("start_time")
	var startTime time.Time
	if startTimeStr != "" {
		start_time, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid start_time format"})
			return
		}
		startTime = start_time
	}

	endTimeStr := ctx.PostForm("end_time")
	var endTime time.Time

	if endTimeStr != "" {
		end_time, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid end_time format"})
			return
		}
		endTime = end_time
	}

	file, err := ctx.FormFile("image")
	var imagePath string
	if err == nil {
		imagePath = fmt.Sprintf("./uploads/%s", file.Filename)

		err = ctx.SaveUploadedFile(file, imagePath)

		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid image format"})
			return
		}
	}

	activity := domain.ActivityUpdate{
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
		Image:     imagePath,
	}
	fmt.Println(activity)

	updatedActivity, err := ac.ActivityUsecase.UpdateActivity(id, activity)

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
