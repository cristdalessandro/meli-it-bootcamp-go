package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Prods Products

func main() {
	var err error
	server := gin.Default()

	err = Prods.load()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Products loaded")

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	prodsRouter := server.Group("/products")
	prodsRouter.GET("/", getAllHandler)
	prodsRouter.GET("/:id", getOneHandler)
	prodsRouter.GET("/search", getPriceGreaterThanHandler)

	server.Run(":3000")

}
