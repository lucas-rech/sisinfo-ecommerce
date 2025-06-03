package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/handler"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils"
)

func SetupRouter(productHandler *handler.ProductHandler, userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{

		admin := v1.Group("/admin")
		{
			// Rotas de produtos
			admin.Use(utils.JWTAuth())
			admin.POST("/product", productHandler.CreateProduct)
			admin.PATCH("/product/:id", productHandler.UpdateProduct)
			admin.DELETE("/product/:id", productHandler.DeleteProduct)

			// Rotas de usuários
			admin.GET("/user/:id", userHandler.FindByID)
			admin.GET("/user/email/:email", userHandler.FindByEmail)
			admin.PATCH("/user/:id", userHandler.UpdateUser)
			admin.DELETE("/user/:id", userHandler.DeleteUser)

		}

		consumer := v1.Group("/")
		{
			consumer.Use(utils.JWTAuthCustomer())
			consumer.GET("/product/:id", productHandler.FindProductByID)
			consumer.GET("/products", productHandler.FindAllProducts)

		}

		// Rotas de autenticação
		v1.POST("/login", userHandler.Login)
		v1.POST("/user", userHandler.CreateUser)
	}

	return router
}
