package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/handler"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils/middleware"
)

func SetupRouter(productHandler *handler.ProductHandler, userHandler *handler.UserHandler, cartItemHandler *handler.CartItemHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	routes := router.Group("/api/v1")
	{

		admin := routes.Group("/admin")
		{
			// Middlewares para autenticação
			admin.Use(middleware.JWTAuthAdmin())

			// Rotas de produtos
			admin.POST("/product", productHandler.CreateProduct)
			admin.PATCH("/product/:id", productHandler.UpdateProduct)
			admin.DELETE("/product/:id", productHandler.DeleteProduct)

			// Rotas de usuários
			admin.GET("/user/:id", userHandler.FindByID)
			admin.GET("/user/email/:email", userHandler.FindByEmail)
			admin.PATCH("/user/:id", userHandler.UpdateUser)
			admin.DELETE("/user/:id", userHandler.DeleteUser)

		}

		consumer := routes.Group("/")
		{
			// Middlewares para autenticação e extração de claims
			consumer.Use(middleware.JWTAuthCustomer(), middleware.ExtractClaims())

			consumer.POST("/cart/item", cartItemHandler.AddItemToCart)
			consumer.PATCH("/cart/item", cartItemHandler.UpdateItemInCart)
			consumer.DELETE("/cart/item", cartItemHandler.RemoveItemFromCart)
			consumer.GET("/cart/items", cartItemHandler.GetAllItemsInCart)

		}

		open := routes.Group("/")
		{
			// Rotas de login e registro de usuários
			open.POST("/login", userHandler.Login)
			open.POST("/login/register", userHandler.CreateUser)


			// Rotas de produtos
			open.GET("/product/:id", productHandler.FindProductByID)
			open.GET("/products", productHandler.FindAllProducts)
		}

	}

	return router
}
