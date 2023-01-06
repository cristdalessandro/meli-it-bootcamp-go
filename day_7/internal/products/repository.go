package products

import (
	types "day_7/pkg"
	"encoding/json"
	"io"
	"os"
)

type Products struct {
	list []Product
	id   int
}

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func (ps *Products) Load() error {
	var (
		productsJson  *os.File
		productsBytes []byte
		err           error
	)

	productsJson, err = os.Open("products.json")
	if err != nil {
		return err
	}
	defer productsJson.Close()

	productsBytes, err = io.ReadAll(productsJson)
	if err != nil {
		return err
	}

	json.Unmarshal(productsBytes, &(ps.list))
	ps.id = 501
	return nil
}

func (ps *Products) Create(data types.Request) {
	newProduct := Product{
		Id:           ps.id,
		Name:         data.Name,
		Quantity:     data.Quantity,
		Code_value:   data.Code_value,
		Is_published: data.Is_published,
		Expiration:   data.Expiration,
		Price:        data.Price,
	}

	ps.list = append(ps.list, newProduct)
	ps.id++
}

func (ps Products) GetAll() []Product {
	return ps.list
}

func (ps Products) GetGreaterThan(min float64) []Product {
	var res []Product
	for _, v := range ps.list {
		if v.Price > min {
			res = append(res, v)
		}
	}
	return res
}

func (ps Products) GetById(id int) *Product {
	for _, v := range ps.list {
		if v.Id == id {
			return &v
		}
	}
	return nil
}

func (ps Products) GetByCodeValue(code string) *Product {
	for _, v := range ps.list {
		if v.Code_value == code {
			return &v
		}
	}
	return nil
}
