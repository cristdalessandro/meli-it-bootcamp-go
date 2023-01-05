package main

import (
	"fmt"
	"strconv"
)

func forecaster() {
	var (
		temperatura int     = 25
		humedad     int     = 50
		presion     float64 = 35.241424
	)
	fmt.Printf("temperatura:%d\nhumedad:%d%%\npresion:%.2f\n", temperatura, humedad, presion)
}

func printName() {
	name, address := "Cristhian", "Montevideo"
	fmt.Printf("Mi nombre es %s y vivo en %s\n", name, address)
}

func spellingWord() {
	var word string
	fmt.Println("Escriba una palabra:")
	_, err := fmt.Scan(&word)
	fmt.Printf("'%s' tiene %d letras\n", word, len(word))
	if err == nil {
		for _, v := range word {
			fmt.Printf("%c\n", v)
		}
	}
}

func loan() {
	var (
		age        int64
		isEmployed bool
		salary     float64
		sage       string
		ssal       string
		sisEmp     string
	)

	fmt.Println("Edad: ")
	_, err1 := fmt.Scan(&sage)
	fmt.Println("Es empleado?(si: 1, no: 0): ")
	_, err2 := fmt.Scan(&sisEmp)
	fmt.Println()
	fmt.Println("Salario: ")
	_, err3 := fmt.Scan(&ssal)
	fmt.Println()

	if err1 != nil || err2 != nil || err3 != nil {
		return
	}

	age, _ = strconv.ParseInt(sage, 10, 64)
	isEmployed = string(sisEmp) == "1"
	salary, _ = strconv.ParseFloat(ssal, 64)

	switch {
	case age < 22:
		{
			fmt.Println("Prestamo denegado, edad minima es 22")
		}
	case !isEmployed:
		{
			fmt.Println("Prestamo denegado, debe tener empleo")
		}
	case salary < 100000:
		{
			fmt.Println("Prestamo denegado, salario minimo es 100000")
		}
	default:
		fmt.Println("Prestamo aprobado")
	}
}
