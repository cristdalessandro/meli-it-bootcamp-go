package main

import (
	"encoding/json"
	"io"
	"os"
)

type Products struct {
	list []Product
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
	return nil
}

func (ps Products) getAll() []Product {
	return ps.list
}

func (ps Products) getGreaterThan(min float64) []Product {
	var res []Product
	for _, v := range ps.list {
		if v.Price > min {
			res = append(res, v)
		}
	}
	return res
}

func (ps Products) getById(id int) *Product {
	for _, v := range ps.list {
		if v.Id == id {
			return &v
		}
	}
	return nil
}
