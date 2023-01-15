package controllers

import (
	"fmt"
	"net/http"

	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/MadMaxMR/Products-Restful/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	page := c.Query("page")

	err := database.GetAll(&products, page)
	if err != nil {
		fmt.Println("error es: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
	product := models.Product{}
	sku := c.Param("sku")

	err := database.Get(&product, sku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	product := models.Product{}

	err := database.Create(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {

}
