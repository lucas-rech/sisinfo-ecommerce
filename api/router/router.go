package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/handler"
)

func SetupRouter(productHandler *handler.ProductHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
	
		// Rotas de produtos
		v1.POST("/product", productHandler.CreateProduct)
		v1.GET("/product/:id", productHandler.FindProductByID)
		v1.GET("/products", productHandler.FindAllProducts)
		v1.PATCH("/product/:id", productHandler.UpdateProduct)
		v1.DELETE("/product/:id", productHandler.DeleteProduct)
		
	}

	return router																				
}
