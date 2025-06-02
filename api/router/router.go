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
	
		v1.POST("/product", productHandler.CreateProduct)
	}

	return router
}
