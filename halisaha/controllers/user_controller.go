package controllers

import (
	"halisaha/database"
	"halisaha/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Eksik veya Hatalı Bilgi Girdiniz"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Sifre Hashlenemedi"})
		return
	}

	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)

	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Bu E-posta adresi zaten kullanılıyor"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanici veri tabanına kaydedilemedi"})

	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Kullanici Veri Tabanina Kaydedildi.."})
}

func LoginUser(ctx *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user models.User

	//json verisini alıyor

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "geçersiz e-posta veya şifre"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "e-posta veya şifre yanlış"})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Giriş Başarılı!"})

}

func GetUser(ctx *gin.Context) {

	var users []models.User

	result := database.DB.Find(&users)

	if result.Error != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return

	}

	ctx.JSON(http.StatusOK, users)

}
