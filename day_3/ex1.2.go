package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	position string
	person   Person
}

func (p Person) printPerson() {
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Birthdate: %s\n", p.DateOfBirth)
}

func (emp Employee) printEmployee() {

	fmt.Println("Personal Data")
	emp.person.printPerson()
	fmt.Println("Employee data")
	fmt.Printf("Employee ID: %d\n", emp.ID)
	fmt.Printf("Position: %s\n", emp.position)
}
