package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/MadMaxMR/Products-Restful/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	db := database.GetConnection()
	defer db.Close()

	err := db.Model(&products).Preload("OtherImages").Find(&products).Error
	fmt.Println("error es: ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
	product := models.Product{}
	sku := c.Param("sku")

	db := database.GetConnection()
	defer db.Close()

	err := db.Model(&product).Preload("OtherImages").First(&product, "sku='"+sku+"'").Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	var oImages []models.OtherImages
	formData, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.SKU = formData.Value["sku"][0]
	product.Name = formData.Value["name"][0]
	product.Brand = formData.Value["brand"][0]
	product.Size = formData.Value["size"][0]
	price, _ := strconv.ParseFloat(formData.Value["price"][0], 32)
	product.Price = float32(price)

	file, err := c.FormFile("image")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": "image required"})
		return
	}
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" && file.Header.Get("Content-Type") != "image/jpg" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato invalido" + file.Header.Get("Content-Type")})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var extension = strings.Split(file.Filename, ".")[1]
	var filename string = "product-" + product.SKU + "." + extension
	err = c.SaveUploadedFile(file, "public/img/product/"+filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error subiendo imagen": err.Error()})
		return
	}
	product.PrincipalImage = "http://localhost:8080/img/product/" + filename

	otherImages := formData.File["otherImages"]

	if len(otherImages) != 0 {
		for i, files := range otherImages {
			newOtherImages := models.OtherImages{}
			fmt.Println()
			var oExtension = strings.Split(file.Filename, ".")[1]
			var oFilename string = "product-" + product.SKU + "-" + strconv.Itoa(i) + "." + oExtension
			err = c.SaveUploadedFile(files, "public/otherImages/"+oFilename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			newOtherImages.ProductSKU = product.SKU
			newOtherImages.ImageURL = "http://localhost:8080/img/product/" + oFilename
			oImages = append(oImages, newOtherImages)
		}
		product.OtherImages = oImages
	}

	err = database.Create(&product)
	if err != nil {
		if strings.Contains(err.Error(), "llave duplicada viola restricción de unicidad") {
			c.JSON(http.StatusConflict, gin.H{"error": "sku duplicado"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	var oImages []models.OtherImages
	sku := c.Param("sku")

	formData, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.SKU = formData.Value["sku"][0]
	product.Name = formData.Value["name"][0]
	product.Brand = formData.Value["brand"][0]
	product.Size = formData.Value["size"][0]
	price, _ := strconv.ParseFloat(formData.Value["price"][0], 32)
	product.Price = float32(price)

	file, err := c.FormFile("image")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": "image required"})
		return
	}
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato invalido de imagen"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var extension = strings.Split(file.Filename, ".")[1]
	var filename string = "product-" + product.SKU + "." + extension
	err = c.SaveUploadedFile(file, "public/img/product/"+filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error subiendo imagen": err.Error()})
		return
	}
	product.PrincipalImage = "http://localhost:8080/img/product/" + filename

	otherImages := formData.File["otherImages"]

	if len(otherImages) != 0 {
		for i, files := range otherImages {
			newOtherImages := models.OtherImages{}
			fmt.Println()
			var oExtension = strings.Split(file.Filename, ".")[1]
			var oFilename string = "product-" + product.SKU + "-" + strconv.Itoa(i) + "." + oExtension
			err = c.SaveUploadedFile(files, "public/otherImages/"+oFilename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			newOtherImages.ProductSKU = product.SKU
			newOtherImages.ImageURL = "http://localhost:8080/img/product/" + oFilename
			oImages = append(oImages, newOtherImages)
		}
		product.OtherImages = oImages
	}

	err = database.Update(&product, "sku='"+sku+"'")
	if err != nil {
		if strings.Contains(err.Error(), "llave duplicada viola restricción de unicidad") {
			c.JSON(http.StatusConflict, gin.H{"error": "sku duplicado"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	sku := c.Param("sku")

	err := database.Delete(&product, "sku='"+sku+"'")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var extension = strings.Split(product.PrincipalImage, ".")[1]
	os.Remove("public/img/product/product-" + sku + "." + extension)

	ruta := filepath.Join("public/otherImages", "product-"+sku+"*")
	fmt.Print("Ruta Impresa= ", ruta)

	image, err := filepath.Glob(ruta)
	fmt.Print("Image Impresa= ", image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, imagen := range image {
		err := os.Remove(imagen)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Eliminado correctamente"})
}

func ViewImage(c *gin.Context) {
	img := c.Param("img")

	_, err := os.Stat("public/img/product/" + img)
	if err == nil {
		c.File("public/img/product/" + img)
		return
	}
	_, err = os.Stat("public/otherImages/" + img)
	if err == nil {
		c.File("public/otherImages/" + img)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"error": "Imagen no encontrada"})
}
