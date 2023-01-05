package main

var Products []Product

type Product struct {
	ID          int
	Name        string
	Price       float64
	description string
	category    string
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() []Product {
	return Products
}

func GetById(id int) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	panic("The product doesn't exist")
}
