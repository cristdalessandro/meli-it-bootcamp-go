package main

import "errors"

var lessThanTenkError error = errors.New("error: el salario es menor a 10.000")

func lessThan10k_b(salary int) error {
	if salary < 10000 {
		return lessThanTenkError
	}
	return nil
}
