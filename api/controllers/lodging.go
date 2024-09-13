package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	usecase_interfaces "github.com/Efamamo/WonderBeam/api/interfaces"
	"github.com/Efamamo/WonderBeam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type LodgingController struct {
	LodgingUsecase usecase_interfaces.ILodging
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
	locationId, err := uuid.Parse(location)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := ctx.PostForm("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	googleLink := ctx.PostForm("google_link")
	if googleLink == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "google_link is required"})
		return
	}

	description := ctx.PostForm("description")
	if description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
		return
	}

	budget_per_night, err := strconv.ParseFloat(ctx.PostForm("budget_per_night"), 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "budget should be number"})
		return
	}

	category := ctx.PostForm("category")
	if category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "category is required"})
		return
	}

	quality_rating, err := strconv.Atoi(ctx.PostForm("quality_rating"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "quality_rating should be number"})
		return
	}

	user_rating, err := strconv.Atoi(ctx.PostForm("user_rating"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user_rating should be number"})
		return
	}

	emails := pq.StringArray(ctx.PostFormArray("emails"))
	if emails == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "emails are required"})
		return
	}

	phone_numbers := pq.StringArray(ctx.PostFormArray("phone_numbers"))
	if phone_numbers == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "phone_numbers are required"})
		return
	}

	website := ctx.PostForm("webite")

	amenities := pq.StringArray(ctx.PostFormArray("amenities"))
	if amenities == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "amenities are required"})
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

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "cant save image"})
		return
	}

	lodging := domain.Lodging{
		Name:           name,
		GoogleLink:     googleLink,
		Description:    description,
		BudgetPerNight: budget_per_night,
		Category:       category,
		QualityRating:  quality_rating,
		UserRating:     user_rating,
		Emails:         emails,
		PhoneNumbers:   phone_numbers,
		Website:        website,
		Amenities:      amenities,
		LocationID:     locationId,
		Image:          imagePath,
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

	name := ctx.PostForm("name")
	googleLink := ctx.PostForm("google_link")
	description := ctx.PostForm("description")

	budget_per_night := ctx.PostForm("budget_per_night")
	var budgetPerNight float64
	if budget_per_night != "" {
		budget, err := strconv.ParseFloat(budget_per_night, 64)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "budget should be number"})
			return
		}
		budgetPerNight = budget
	}
	category := ctx.PostForm("category")

	quality_rating := ctx.PostForm("quality_rating")
	var qualityRating int

	if quality_rating != "" {
		quality, err := strconv.Atoi(quality_rating)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "quality_rating should be number"})
			return
		}
		qualityRating = quality
	}

	user_rating := ctx.PostForm("user_rating")
	var userRating int

	if user_rating != "" {
		user, err := strconv.Atoi(user_rating)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user_rating should be number"})
			return
		}
		qualityRating = user
	}

	emails := pq.StringArray(ctx.PostFormArray("emails"))
	phone_numbers := pq.StringArray(ctx.PostFormArray("phone_numbers"))
	website := ctx.PostForm("webite")
	amenities := pq.StringArray(ctx.PostFormArray("amenities"))

	file, err := ctx.FormFile("image")
	var imagePath string
	if err == nil {
		imagePath = fmt.Sprintf("./uploads/%s", file.Filename)
		err = ctx.SaveUploadedFile(file, imagePath)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "cant save image"})
			return
		}
	}

	lodging := domain.LodgingUpdate{
		Name:           name,
		GoogleLink:     googleLink,
		Description:    description,
		BudgetPerNight: budgetPerNight,
		Category:       category,
		QualityRating:  qualityRating,
		UserRating:     userRating,
		Emails:         emails,
		PhoneNumbers:   phone_numbers,
		Website:        website,
		Amenities:      amenities,
		Image:          imagePath,
	}
	updatedLodging, err := ldc.LodgingUsecase.UpdateLodging(id, lodging)

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
