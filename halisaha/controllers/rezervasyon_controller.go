package controllers

import (
	"halisaha/database"
	"halisaha/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// REZERVASYON LİSTELEME
func GetRezervations(ctx *gin.Context) {

	var rezervasyonlar []models.Rezervasyon

	result := database.DB.Find(&rezervasyonlar)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return

	}

	ctx.JSON(http.StatusOK, rezervasyonlar)

}

func CreateRezervation(ctx *gin.Context) {

	var rezervasyon models.Rezervasyon

	if err := ctx.ShouldBindJSON(&rezervasyon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Eksik veya Hatali veri girdiniz"})
		return
	}

	var existingRezervasyon models.Rezervasyon
	if err := database.DB.Where("saha_id=? AND tarih=? AND saat=?", rezervasyon.SahaID, rezervasyon.Tarih, rezervasyon.Saat).First(&existingRezervasyon).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "bu saat için zaten rezervasyon bulunuyor"})
		return
	}

	result := database.DB.Create(&rezervasyon)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Rezervasyon Oluşturulamadı"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Rezervasyon Başarıyla Oluşturuldu"})
}
