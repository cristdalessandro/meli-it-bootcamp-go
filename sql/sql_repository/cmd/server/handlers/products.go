package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/domain"
	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/products"
	rb "github.com/cristdalessandro/meli-it-bootcamp-go/sql/pkg/requestBody"

	"github.com/gin-gonic/gin"
)

type productsHandler struct {
	s products.Services
}

func NewHandler(s products.Services) *productsHandler {
	return &productsHandler{s: s}
}

func (handlers *productsHandler) GetAll(ctx *gin.Context) {
	products, err := handlers.s.GetAll()
	if err != nil {
		ctx.JSON(500, err.Error())
	}

	ctx.JSON(200, products)
}

func (handlers *productsHandler) GetOne(ctx *gin.Context) {
	var (
		err     error
		id      int
		product domain.Product
	)
	idParam := ctx.Param("id")

	id, err = strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	product, err = handlers.s.GetOne(id)
	if err != nil {
		ctx.AbortWithError(404, errors.New("product not found"))
		return
	}

	ctx.JSON(200, product)
}

func (handlers *productsHandler) Create(ctx *gin.Context) {
	var (
		req rb.ProductsPost
		err error
	)

	if err = ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if _, err := time.Parse("2006-01-02", req.Expiration); err != nil {
		ctx.AbortWithError(400, errors.New("invalid date format"))
		return
	}

	p := domain.Product{
		Name:         req.Name,
		Quantity:     req.Quantity,
		Code_value:   req.Code_value,
		Is_published: req.Is_published,
		Expiration:   req.Expiration,
		Price:        req.Price,
	}

	res, err := handlers.s.Create(p)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(201, res)
}

// func EditHandler(ctx *gin.Context) {
// 	var (
// 		id  int
// 		req rb.ProductsPatch
// 		err error
// 	)
// 	idParam := ctx.Param("id")

// 	id, err = strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.AbortWithError(400, err)
// 		return
// 	}

// 	if err = ctx.ShouldBind(&req); err != nil {
// 		ctx.AbortWithError(400, err)
// 		return
// 	}

// 	s, err := productsServices.GetProductServicesInstance()

// 	if err != nil {
// 		ctx.AbortWithError(500, err)
// 	}

// 	err = s.EditProduct(req, id)

// 	if err != nil {
// 		ctx.AbortWithError(400, errors.New("bad request"))
// 		return
// 	}

// }

// func DeleteHandler(ctx *gin.Context) {
// 	var id int
// 	idParam := ctx.Param("id")
// 	s, err := productsServices.GetProductServicesInstance()

// 	if err != nil {
// 		ctx.AbortWithError(500, err)
// 	}

// 	id, err = strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.AbortWithError(400, err)
// 		return
// 	}

// 	err = s.DeleteProduct(id)
// 	if err != nil {
// 		ctx.AbortWithError(404, errors.New("product not found"))
// 		return
// 	}

// 	ctx.JSON(204, "")
// }
