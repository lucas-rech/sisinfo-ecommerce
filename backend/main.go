package main

import (
	"log"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/router"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/config"


	_ "github.com/lucas-rech/sisinfo-ecommerce/backend/docs" 
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sisinfo E-commerce API
// @version 1.0
// @description This is a sample e-commerce API for Sisinfo course.
// @termsOfService http://swagger.io/terms/
// @basePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization


func main() {
	container, err := config.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	if sqlDB, err := container.DB.DB(); err == nil {
		defer func() {
			log.Println("Closing database connection")
			if closeErr := sqlDB.Close(); closeErr != nil {
				log.Fatalf("Failed to close database connection: %v", closeErr)
			}
		}()
	} else {
		log.Printf("Failed to get database connection: %v", err)
	}


	r := router.SetupRouter(container.ProductHandler, container.UserHandler, container.CartItemHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}