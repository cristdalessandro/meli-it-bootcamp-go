package main

import (
	"fmt"

	tickets "github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	var total, tByP, avgDest int
	var err error

	tickets.ParseCSV()
	total, err = tickets.GetTotalTickets("Argentina")
	if err != nil {
		panic("Error calculando tickets totales")
	}
	fmt.Printf("Brasil total: %d\n", total)

	tByP, err = tickets.GetCountByPeriod("tarde")
	if err != nil {
		panic("Error calculando tickets por periodo")
	}
	fmt.Printf("Tarde: %d\n", tByP)

	avgDest, err = tickets.AverageDestination("Indonesia")
	if err != nil {
		panic("Error calculando el promedio por destino")
	}
	fmt.Printf("Indonesia average: %d%%\n", avgDest)
}
