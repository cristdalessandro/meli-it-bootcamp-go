package main

import (
	"fmt"
	"os"
)

func readFileCustomer() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ejecución finalizada")
	}()

	_, err := os.ReadFile("customerss.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}
