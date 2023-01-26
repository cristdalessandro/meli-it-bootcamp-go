package products

import (
	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/domain"
)

type Services interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
}

type services struct {
	repo Repository
}

func NewServices(repo Repository) Services {
	return &services{repo: repo}
}

func (s *services) GetAll() (res []domain.Product, err error) {
	return s.repo.GetAll()
}

func (s *services) GetOne(id int) (res domain.Product, err error) {
	return s.repo.GetOne(id)
}

func (s *services) Create(p domain.Product) (res domain.Product, err error) {
	return s.repo.Create(p)
}
