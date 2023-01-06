package handlers

import (
	productsServices "day_7/internal/products"
	types "day_7/pkg"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var Prods productsServices.Products

func CreateProduct(ctx *gin.Context) {

	var req types.Request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if product := Prods.GetByCodeValue(req.Code_value); product != nil {
		ctx.AbortWithError(400, errors.New("product already exist"))
		return
	}

	if _, err := time.Parse("02/01/2006", req.Expiration); err != nil {
		ctx.AbortWithError(400, errors.New("invalid date format"))
		return
	}

	Prods.Create(req)

	ctx.JSON(201, "product succesfully created")
}

func GetAllHandler(ctx *gin.Context) {
	products := Prods.GetAll()
	ctx.JSON(200, products)
}

func GetPriceGreaterThanHandler(ctx *gin.Context) {
	min, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	products := Prods.GetGreaterThan(min)

	ctx.JSON(200, products)
}

func GetOneHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	product := Prods.GetById(id)
	if product == nil {
		ctx.AbortWithError(404, errors.New("product not found"))
		return
	}
	ctx.JSON(200, product)
}
