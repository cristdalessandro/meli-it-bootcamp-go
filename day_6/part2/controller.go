package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllHandler(ctx *gin.Context) {
	products := Prods.getAll()
	ctx.JSON(200, products)
}

func getPriceGreaterThanHandler(ctx *gin.Context) {
	min, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	products := Prods.getGreaterThan(min)

	ctx.JSON(200, products)
}

func getOneHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	product := Prods.getById(id)
	ctx.JSON(200, product)
}
