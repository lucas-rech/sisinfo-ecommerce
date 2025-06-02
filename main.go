package main

import (
	"log"

	"github.com/lucas-rech/sisinfo-ecommerce/api/router"
	"github.com/lucas-rech/sisinfo-ecommerce/config"


	_ "github.com/lucas-rech/sisinfo-ecommerce/docs" // This line is necessary for go-swagger to generate the docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sisinfo E-commerce API
// @version 1.0
// @description This is a sample e-commerce API for Sisinfo course.
// @termsOfService http://swagger.io/terms/
// @basePath /api/v1


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


	r := router.SetupRouter(container.ProductHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}