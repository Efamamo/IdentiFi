package controllers

import (
	"fmt"
	"net/http"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
)

type LocationController struct {
	LocationUseCase usecase_interfaces.ILocation
}

func (lc LocationController) AddLocation(ctx *gin.Context) {
	name := ctx.PostForm("name")
	googleLink := ctx.PostForm("google_link")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if googleLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "google_ink is required"})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	imagePath := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
		return
	}

	location := domain.Location{
		Name:       name,
		GoogleLink: googleLink,
		Image:      imagePath,
	}

	loc, err := lc.LocationUseCase.AddLocation(location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, *loc)
}

func (lc LocationController) UpdateLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.PostForm("name")
	googleLink := ctx.PostForm("google_link")

	file, err := ctx.FormFile("image")
	var imagePath string
	if err == nil {
		imagePath = fmt.Sprintf("./uploads/%s", file.Filename)
		if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}
	}

	location := domain.LocationUpdate{
		Name:       name,
		GoogleLink: googleLink,
		Image:      imagePath,
	}

	loc, err := lc.LocationUseCase.UpdateLocation(id, location)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, *loc)
}

func (lc LocationController) GetLocations(ctx *gin.Context) {

	locations, err := lc.LocationUseCase.GetLocations()

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted, locations)
}


func (lc LocationController) GetLocationById(ctx *gin.Context) {

	location, err := lc.LocationUseCase.GetLocationById(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted, location)
}

func (lc LocationController) DeleteLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	err := lc.LocationUseCase.DeleteLocation(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	ctx.Status(http.StatusNoContent)
}
