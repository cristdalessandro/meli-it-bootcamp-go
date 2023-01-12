package products

import (
	types "github.com/cristdalessandro/meli-it-bootcamp-go/day_9/pkg"
)

type ProductServices struct {
	repository IProducts
}

var services *ProductServices

func GetProductServicesInstance() (*ProductServices, error) {
	if services == nil {
		services = &ProductServices{
			repository: &Products{},
		}

		err := services.repository.load()
		if err != nil {
			return nil, err
		}
	}
	return services, nil
}

func (services *ProductServices) CreateProduct(data types.Request) {
	services.repository.create(data)
}

func (services *ProductServices) ReplaceProduct(data types.Request, id int) error {
	return services.repository.replace(data, id)
}

func (services *ProductServices) EditProduct(data types.PatchRequest, id int) error {
	return services.repository.edit(data, id)
}

func (services *ProductServices) DeleteProduct(id int) error {
	return services.repository.delete(id)
}

func (services *ProductServices) GetAllProducts() []Product {
	return services.repository.getAll()
}

func (services *ProductServices) GetProductsPricedMoreThan(min float64) []Product {
	return services.repository.getGreaterThan(min)
}

func (services *ProductServices) GetProductById(id int) *Product {
	return services.repository.getById(id)
}

func (services *ProductServices) GetProductByCodeValue(code string) *Product {
	return services.repository.getByCodeValue(code)
}
