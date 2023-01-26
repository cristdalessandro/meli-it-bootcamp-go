package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/cmd/server/handlers"
	products "github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/products"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.rg = r.eng.Group("/api/v1")

	r.setProductsRoutes()

}

func (r *router) setProductsRoutes() {
	pr := r.rg.Group("/products")

	repo := products.NewRepository(r.db)
	services := products.NewServices(repo)
	handlers := handlers.NewHandler(services)

	pr.GET("/", handlers.GetAll)
	pr.GET("/:id", handlers.GetOne)
	pr.POST("/", handlers.Create)
	// pr.PATCH("/:id", pc.EditHandler)
	// pr.DELETE("/:id", pc.DeleteHandler)

}
