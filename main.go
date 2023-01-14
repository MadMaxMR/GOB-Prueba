package main

import (
	"os"

	"github.com/MadMaxMR/product-rest/controllers"
	"github.com/MadMaxMR/product-rest/database"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = defaultPort
	}

	//db := database.GetConnection()
	database.Migrate()

	r := gin.Default()

	r.GET("/product", controllers.GetProducts)
	r.GET("/product/:sku", controllers.GetProduct)
	r.POST("/saveProduct", controllers.CreateProduct)

	r.Run(":" + serverPort)
}
