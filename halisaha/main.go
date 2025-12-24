package main

import (
	"fmt"
	"halisaha/database"
	"halisaha/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	port := ":8080"

	fmt.Println("**Sunucu Baslatiliyor** http://localhost" + port)

	r.Run(port)
}
