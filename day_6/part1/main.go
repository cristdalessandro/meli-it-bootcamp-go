package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	type fullName struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	// exercise 1
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	// exercise 2
	router.POST("/saludo", func(ctx *gin.Context) {
		body := fullName{}

		if err := ctx.ShouldBind(&body); err != nil {
			ctx.AbortWithStatus(400)
			return
		}

		ctx.String(200, "Hola "+body.FirstName+" "+body.LastName)
	})

	router.Run(":3000")
}
