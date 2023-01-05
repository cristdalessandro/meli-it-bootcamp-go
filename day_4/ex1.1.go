package main

type CustomError150 struct {
}

func (customError CustomError150) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func needPay(salary int) error {
	if salary < 150000 {
		return CustomError150{}
	}
	return nil
}
