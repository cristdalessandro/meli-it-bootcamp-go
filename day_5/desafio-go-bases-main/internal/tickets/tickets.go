package tickets

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

type Ticket struct {
	nombre      string
	email       string
	paisDestino string
	horaVuelo   time.Time
	precio      float64
}

type partsOfDayCounter struct {
	madrugada int
	maniana   int
	tarde     int
	noche     int
}

var tickets []Ticket

func ParseCSV() {
	f, err := os.Open("tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(tickets) > 0 {
		return
	}

	for _, line := range data {
		var t Ticket
		for j, field := range line {
			switch j {
			case 1:
				{
					t.nombre = field
				}
			case 2:
				{
					t.email = field
				}
			case 3:
				{
					t.paisDestino = field
				}
			case 4:
				{
					v, _ := time.Parse("15:04", field)
					t.horaVuelo = v
				}
			case 5:
				{
					v, _ := strconv.ParseFloat(field, 64)
					t.precio = v
				}
			}
		}
		tickets = append(tickets, t)
	}
}

// ejercicio 1
func GetTotalTickets(destination string) (res int, err error) {
	for _, v := range tickets {
		if v.paisDestino == destination {
			res++
		}
	}
	return
}

// ejercicio 2
func countTotalByPartOfDay() (res partsOfDayCounter, err error) {
	for _, v := range tickets {
		switch {
		case v.horaVuelo.Hour() < 0 || v.horaVuelo.Hour() > 23:
			{
				err = errors.New("hora invalida")
			}
		case v.horaVuelo.Hour() < 7:
			{
				res.madrugada++
			}
		case v.horaVuelo.Hour() < 13:
			{
				res.maniana++
			}
		case v.horaVuelo.Hour() < 20:
			{
				res.tarde++
			}
		case v.horaVuelo.Hour() < 24:
			{
				res.noche++
			}
		}
	}
	return
}

func GetCountByPeriod(time string) (a int, err error) {
	counters, e := countTotalByPartOfDay()
	if err != nil {
		err = e
		return
	}
	switch time {
	case "madrugada":
		{
			a = counters.madrugada
		}
	case "mañana":
		{
			a = counters.maniana
		}
	case "tarde":
		{
			a = counters.tarde
		}
	case "noche":
		{
			a = counters.noche
		}
	}
	return
}

// ejercicio 3
func AverageDestination(destination string) (avg int, err error) {
	totalByCountry, _ := GetTotalTickets(destination)
	avg = 100 * totalByCountry / len(tickets)
	return
}
