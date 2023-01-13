package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	middles "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/cmd/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("SECRET", "12345")
	r := gin.New()

	rp := r.Group("/products")
	rp.Use(middles.CheckToken)

	rp.POST("/", CreateProduct)
	rp.DELETE("/:id", DeleteHandler)
	rp.PUT("/:id", ReplaceHandler)
	rp.PATCH("/:id", EditHandler)
	rp.GET("/", GetAllHandler)
	rp.GET("/:id", GetOneHandler)
	rp.GET("/search", GetPriceGreaterThanHandler)

	return r
}

func createRequest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 12345")

	return req, httptest.NewRecorder()
}

func TestGetAll(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/products/", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestGetOne(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/products/1", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestCreate(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodPost, "/products/", `{
		"name": "Beer",
		"quantity": 100,
		"code_value": "SIS",
		"expiration": "30/08/2020",
		"price": 1100
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 201, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodDelete, "/products/1", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 204, rr.Code)
}

func TestInvalidIdParam(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/products/hello", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Code)
}

func TestNotFound(t *testing.T) {
	r := createServer()

	req, rr := createRequest(http.MethodGet, "/products/700", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)
}

func TestUnauthorized(t *testing.T) {
	r := createServer()
	os.Setenv("SECRET", "ABC")

	req, rr := createRequest(http.MethodGet, "/products/hello", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 401, rr.Code)
}
