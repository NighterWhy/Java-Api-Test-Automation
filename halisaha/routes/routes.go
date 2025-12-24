package routes

import (
	"halisaha/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")

	//Kullanici Islemleri
	{
		userRoutes.GET("/", controllers.GetUser)
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}

	//Hali saha İşlemleri
	sahaRoutes := r.Group("/sahalar")

	{
		sahaRoutes.GET("/", controllers.GetSahalar)
		sahaRoutes.POST("/", controllers.CreateSaha)
	}

	//rezervasyo işlemleri

	rezervasyonRoutes := r.Group("/rezervasyonlar")

	{
		rezervasyonRoutes.GET("/", controllers.GetRezervations)
		rezervasyonRoutes.POST("/", controllers.CreateRezervation)
	}
}
