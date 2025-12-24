package controllers

import (
	"halisaha/database"
	"halisaha/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSahalar(ctx *gin.Context) {

	var sahalar []models.Saha

	result := database.DB.Find(&sahalar)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sahalar)

}

//Halı Saha Ekleme Fonksiyonu

func CreateSaha(ctx *gin.Context) {
	var saha models.Saha

	if err := ctx.ShouldBindJSON(&saha); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	result := database.DB.Create(&saha)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, saha)
}
