package handlers

import (
	"errors"
	"strconv"
	"time"

	productsServices "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/internal/products"
	types "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/pkg"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var req types.Request
	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if product := s.GetProductByCodeValue(req.Code_value); product != nil {
		ctx.AbortWithError(400, errors.New("product already exist"))
		return
	}

	if _, err := time.Parse("02/01/2006", req.Expiration); err != nil {
		ctx.AbortWithError(400, errors.New("invalid date format"))
		return
	}

	s.CreateProduct(req)

	ctx.JSON(201, "product succesfully created")
}

func ReplaceHandler(ctx *gin.Context) {
	var (
		id  int
		req types.Request
		err error
	)
	idParam := ctx.Param("id")

	id, err = strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	err = s.ReplaceProduct(req, id)

	if err != nil {
		ctx.AbortWithError(400, errors.New("bad request"))
		return
	}

	ctx.JSON(204, "")
}

func EditHandler(ctx *gin.Context) {
	var (
		id  int
		req types.PatchRequest
		err error
	)
	idParam := ctx.Param("id")

	id, err = strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	err = s.EditProduct(req, id)

	if err != nil {
		ctx.AbortWithError(400, errors.New("bad request"))
		return
	}

}

func DeleteHandler(ctx *gin.Context) {
	var id int
	idParam := ctx.Param("id")
	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	id, err = strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = s.DeleteProduct(id)
	if err != nil {
		ctx.AbortWithError(404, errors.New("product not found"))
		return
	}

	ctx.JSON(204, "")
}

func GetAllHandler(ctx *gin.Context) {
	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	products := s.GetAllProducts()
	ctx.JSON(200, products)
}

func GetPriceGreaterThanHandler(ctx *gin.Context) {
	var s *productsServices.ProductServices
	min, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	s, err = productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	products := s.GetProductsPricedMoreThan(min)

	ctx.JSON(200, products)
}

func GetOneHandler(ctx *gin.Context) {
	var id int
	idParam := ctx.Param("id")
	s, err := productsServices.GetProductServicesInstance()

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	id, err = strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	product := s.GetProductById(id)
	if product == nil {
		ctx.AbortWithError(404, errors.New("product not found"))
		return
	}

	ctx.JSON(200, product)
}
