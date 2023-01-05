package main

import "errors"

func salaryCalculation(hrs float64, value float64) (salary float64, err error) {
	salary = hrs * value
	if hrs < 80 {
		err = errors.New("Error: el trabajador no pudo haber trabajado menos de 80 horas mensuales")
		return
	}
	if salary > 150000 {
		salary *= 0.9
	}
	return

}
