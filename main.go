package main

import (
	"log"
	"os"

	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/MadMaxMR/Products-Restful/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = defaultPort
	}

	database.Migrate()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))

	routes.SetProductsRoutes(router)

	err := router.Run(":" + serverPort)
	if err != nil {
		log.Fatal("Error al iniciar Servidor")
		return
	}
}
