package routes

import (
	productsController "github.com/cristdalessandro/meli-it-bootcamp-go/day_8/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(baseUrl string, server *gin.Engine) {
	productsRouter := server.Group("/products")

	productsRouter.POST("/", productsController.CreateProduct)
	productsRouter.DELETE("/:id", productsController.DeleteHandler)
	productsRouter.PUT("/:id", productsController.ReplaceHandler)
	productsRouter.PATCH("/:id", productsController.EditHandler)
	productsRouter.GET("/", productsController.GetAllHandler)
	productsRouter.GET("/:id", productsController.GetOneHandler)
	productsRouter.GET("/search", productsController.GetPriceGreaterThanHandler)

}
