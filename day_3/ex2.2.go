package main

const (
	s = "small"
	m = "medium"
	l = "large"
)

type IProduct interface {
	Price() float64
}

type Small struct {
	basePrice float64
}

func (p *Small) Price() float64 {
	return p.basePrice
}

type Medium struct {
	basePrice float64
}

func (p *Medium) Price() float64 {
	return p.basePrice * 1.3
}

type Large struct {
	basePrice float64
}

func (p *Large) Price() float64 {
	return p.basePrice*1.6 + 2500
}

func factory(price float64, pType string) IProduct {
	var product IProduct
	switch pType {
	case s:
		product = &Small{
			basePrice: price,
		}
	case m:
		product = &Medium{
			basePrice: price,
		}
	case l:
		product = &Large{
			basePrice: price,
		}
	}
	return product
}
