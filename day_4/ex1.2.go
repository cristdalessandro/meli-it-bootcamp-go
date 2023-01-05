package main

type CustomError10 struct {
}

func (customError CustomError10) Error() string {
	return "Error: el salario es menor a 10.000"
}

func lessThan10k(salary int) error {
	if salary < 10000 {
		return CustomError10{}
	}
	return nil
}
