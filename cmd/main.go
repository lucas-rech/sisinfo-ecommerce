package main

import "github.com/lucas-rech/sisinfo-ecommerce/api/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}