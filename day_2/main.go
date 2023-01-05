package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	// var input, oper string
	// var nums []float64
	// fmt.Print("Ingrese los numeros separados por espacios: ")
	// reader := bufio.NewReader(os.Stdin)

	// // ReadString will block until the delimiter is entered
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }

	// // remove the delimeter from the string
	// input = strings.TrimSuffix(input, "\n")

	// stringSlice := strings.Split(input, " ")
	// for _, s := range stringSlice {
	// 	n, _ := strconv.ParseFloat(s, 64)
	// 	nums = append(nums, n)
	// }

	// fmt.Print("Ingrese la operacion(minimum, average o maximum): ")
	// fmt.Scan(&oper)
	// op, _ := operation(oper)
	// res := op(nums...)
	// fmt.Println(res)

	fmt.Println("impuestos: ", calculateTaxes(30000))
	fmt.Println("impuestos: ", calculateTaxes(75000))
	fmt.Println("impuestos: ", calculateTaxes(205000))

	fmt.Println("salario: ", salaryCalculation(180, 'C'))
	fmt.Println("salario: ", salaryCalculation(180, 'B'))
	fmt.Println("salario: ", salaryCalculation(180, 'A'))

	fmt.Println("promedio de notas: ", gradesAverage(6, 10, 4, 4, 6))

	min, ok := operation("minimum")
	if !ok {
		return
	}
	ave, ok := operation("average")
	if !ok {
		return
	}
	max, ok := operation("maximum")
	if !ok {
		return
	}
	fmt.Println(min(2, 43, 64, 3, 1, 3))
	fmt.Println(ave(2, 43, 64, 3, 1, 3))
	fmt.Println(max(2, 43, 64, 3, 1, 3))
	_, ok = operation("invalido")
	if !ok {
		fmt.Println("error")
	}

	kgDog := animal("perro")
	kgCat := animal("gato")
	kgSpider := animal("tarantula")
	kgHamster := animal("hamster")
	fmt.Printf("total kilos para perro: %.2f\n", kgDog(10))
	fmt.Printf("total kilos para gato: %.2f\n", kgCat(10))
	fmt.Printf("total kilos para hamster: %.2f\n", kgHamster(10))
	fmt.Printf("total kilos para tarantula: %.2f\n", kgSpider(10))

}
