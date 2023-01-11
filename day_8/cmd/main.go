package main

import (
	productsRouter "github.com/cristdalessandro/meli-it-bootcamp-go/day_8/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	productsRouter.Setup("/products", server)

	server.Run(":3000")
}
