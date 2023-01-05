package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Ejercicio 1.1")
	fmt.Println()

	p1 := Product{
		ID:          1,
		Name:        "phone",
		description: "",
		category:    "electronics",
		Price:       100,
	}
	p2 := Product{
		ID:          2,
		Name:        "watch",
		description: "",
		category:    "electronics",
		Price:       80,
	}

	p1.Save()
	p2.Save()

	fmt.Printf("Products: %v\n", p1.GetAll())
	fmt.Printf("Product 2: %v\n", *GetById(2))

	fmt.Println()
	fmt.Println("Ejercicio 1.2")
	fmt.Println()

	newEmployee := Employee{
		person: Person{
			ID:          1,
			DateOfBirth: "09/02/1967",
			Name:        "Roberto",
		},
		ID:       14,
		position: "Tech lead",
	}

	newEmployee.printEmployee()

	fmt.Println()
	fmt.Println("Ejercicio 2.1")
	fmt.Println()

	date, _ := time.Parse("2006-01-02", "2022-12-19")
	alumnos.add(Alumno{Nombre: "Cristhian", Apellido: "D'Alessandro", DNI: "1.111.111-1", Fecha: date})
	alumnos.add(Alumno{Nombre: "Juan", Apellido: "Ospina", DNI: "2.222.222-2", Fecha: date})

	alumnos.Detalle()

	fmt.Println()
	fmt.Println("Ejercicio 2.2")
	fmt.Println()

	smallProduct := factory(10000, "small")
	mediumProduct := factory(10000, "medium")
	largeProduct := factory(10000, "large")

	fmt.Printf("small(10000): %.2f\n", smallProduct.Price())
	fmt.Printf("medium(10000): %.2f\n", mediumProduct.Price())
	fmt.Printf("large(10000): %.2f\n", largeProduct.Price())

}
