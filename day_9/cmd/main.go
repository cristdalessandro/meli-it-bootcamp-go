package main

import (
	"os"

	productsRouter "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/cmd/routes"
	"github.com/cristdalessandro/meli-it-bootcamp-go/day_9/docs"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI bootcamp
// @version 1.0
// @description This is the go web module practice
// @host localhost:3000
func main() {

	err := godotenv.Load()
	if err != nil {
		panic("error loading environment variables")
	}

	server := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productsRouter.Setup(server)

	server.Run(":3000")
}
