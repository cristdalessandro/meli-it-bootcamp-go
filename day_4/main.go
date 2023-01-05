package main

func main() {

	// var (
	// 	salary int
	// 	err    error
	// )

	// 1.1
	// salary = 80000
	// err = needPay(salary)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Debe pagar impuesto")

	// 1.2
	// salary = 5000
	// err = lessThan10k(salary)
	// if errors.Is(err, CustomError10{}) {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Salario OK")

	// 1.3
	// salary = 5000
	// err = lessThan10k_b(salary)
	// if errors.Is(err, lessThanTenkError) {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Salario OK")

	// 1.4
	// salary = 5000
	// err = lessThan10k_c(salary)
	// if errors.Is(err, lessThanTenkErrorFmt) {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("Salario OK")

	// 1.5
	// hrs, value := 1000.0, 55.0
	// salary, err := salaryCalculation(hrs, value)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Horas: %.2f, Valor Hora: %.2f, Salario: %.2f\n", hrs, value, salary)

	//2.1 & 2.2
	// readFileCustomer()

	//2.3
	user := &User{
		"Contract",
		"Name",
		"DNI",
		"phone",
		"address",
	}

	user2 := &User{
		"Contract",
		"Name",
		"DNI2",
		"",
		"address",
	}

	user.signUp()

	user.signUp()

	user2.signUp()

}
