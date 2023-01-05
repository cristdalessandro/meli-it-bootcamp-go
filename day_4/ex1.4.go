package main

import (
	"errors"
	"fmt"
)

var lessThanTenkErrorFmt = errors.New("error: ")

func lessThan10k_c(salary int) error {
	if salary < 10000 {
		return fmt.Errorf("%w el salario es %d, pero debe ser mayor a 10000", lessThanTenkErrorFmt, salary)
	}
	return nil
}
