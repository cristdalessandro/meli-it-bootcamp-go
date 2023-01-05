package main

const (
	dog     = "perro"
	cat     = "gato"
	spider  = "tarantula"
	hamster = "hamster"
)

type foodKg func(int) float64

func animal(tanimal string) foodKg {
	switch tanimal {
	case dog:
		{
			return animalDog
		}
	case cat:
		{
			return animalCat
		}
	case spider:
		{
			return animalSpider
		}
	case hamster:
		{
			return animalHamster
		}
	}
	panic("invalid animal")
}

func animalDog(amount int) float64 {
	return 10 * float64(amount)
}

func animalCat(amount int) float64 {
	return 5 * float64(amount)
}

func animalSpider(amount int) float64 {
	return 0.15 * float64(amount)
}

func animalHamster(amount int) float64 {
	return 0.25 * float64(amount)
}
