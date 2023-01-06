package main

import (
	productsController "day_7/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	err := productsController.Prods.Load()
	if err != nil {
		panic("data couldn't load")
	}

	server := gin.Default()
	productsRouter := server.Group("/products")

	productsRouter.POST("/", productsController.CreateProduct)
	productsRouter.GET("/", productsController.GetAllHandler)
	productsRouter.GET("/:id", productsController.GetOneHandler)
	productsRouter.GET("/search", productsController.GetPriceGreaterThanHandler)

	server.Run(":3000")
}
