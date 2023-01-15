package main

import (
	"os"

	"github.com/MadMaxMR/Products-Restful/controllers"
	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = defaultPort
	}

	database.Migrate()

	r := gin.Default()

	r.GET("/product", controllers.GetProducts)
	r.GET("/product/:sku", controllers.GetProduct)
	r.POST("/saveProduct", controllers.CreateProduct)

	r.Run(":" + serverPort)
}
