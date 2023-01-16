package routes

import (
	"github.com/MadMaxMR/Products-Restful/controllers"
	"github.com/gin-gonic/gin"
)

func SetProductsRoutes(r *gin.Engine) {

	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:sku", controllers.GetProduct)
	r.PUT("/products/:sku", controllers.UpdateProduct)
	r.DELETE("/products/:sku", controllers.DeleteProduct)

	r.GET("/img/product/:img", controllers.ViewImage)
}
