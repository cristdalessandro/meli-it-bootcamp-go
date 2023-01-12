package main

import (
	productsRouter "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	productsRouter.Setup(server)

	server.Run(":3000")
}
