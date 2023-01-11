package products

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	types "github.com/cristdalessandro/meli-it-bootcamp-go/day_8/pkg"
)

type Products struct {
	list []Product
	id   int
}

type IProducts interface {
	load() error
	create(data types.Request)
	replace(data types.Request, id int) error
	edit(data types.PatchRequest, id int) error
	delete(id int) error
	getAll() []Product
	getGreaterThan(min float64) []Product
	getById(id int) *Product
	getByCodeValue(code string) *Product
}

func (ps *Products) load() error {
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

func (ps *Products) create(data types.Request) {
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

func (ps *Products) replace(data types.Request, id int) (err error) {
	oldProduct := ps.getById(id)

	if oldProduct == nil {
		err = errors.New("not found")
		return
	}

	p := ps.getByCodeValue(data.Code_value)

	if p != nil {
		return errors.New("product already exist")
	}

	oldProduct.Name = data.Name
	oldProduct.Quantity = data.Quantity
	oldProduct.Code_value = data.Code_value
	oldProduct.Is_published = data.Is_published
	oldProduct.Expiration = data.Expiration
	oldProduct.Price = data.Price
	return
}

func (ps *Products) edit(data types.PatchRequest, id int) (err error) {
	oldProduct := ps.getById(id)

	if oldProduct == nil {
		err = errors.New("not found")
		return
	}

	if data.Code_value != nil {
		p := ps.getByCodeValue(*data.Code_value)

		if p != nil {
			return errors.New("product already exist")
		}
		oldProduct.Code_value = *data.Code_value
	}

	if data.Expiration != nil {
		oldProduct.Expiration = *data.Expiration
	}

	if data.Name != nil {
		oldProduct.Name = *data.Name
	}

	if data.Is_published != nil {
		oldProduct.Is_published = *data.Is_published
	}

	if data.Price != nil {
		oldProduct.Price = *data.Price
	}

	if data.Quantity != nil {
		oldProduct.Quantity = *data.Quantity
	}
	return
}

func (ps *Products) delete(id int) (err error) {
	var (
		i int
		v Product
	)

	for i, v = range ps.list {
		if v.Id == id {
			break
		}
	}

	if i == len(ps.list)-1 && ps.list[len(ps.list)-1].Id != id {
		err = errors.New("product not found")
		return
	}

	ps.list = append(ps.list[:i], ps.list[i+1:]...)

	return
}

func (ps *Products) getAll() []Product {
	return ps.list
}

func (ps *Products) getGreaterThan(min float64) []Product {
	var res []Product
	for _, v := range ps.list {
		if v.Price > min {
			res = append(res, v)
		}
	}
	return res
}

func (ps *Products) getById(id int) *Product {
	var (
		i int
		v Product
	)
	for i, v = range ps.list {
		if v.Id == id {
			return &ps.list[i]
		}
	}
	return nil
}

func (ps *Products) getByCodeValue(code string) *Product {
	for _, v := range ps.list {
		if v.Code_value == code {
			return &v
		}
	}
	return nil
}
