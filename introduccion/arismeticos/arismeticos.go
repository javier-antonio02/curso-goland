package main

import "fmt"

func main() {
	a := 20
	b := 10

	//suma
	result := a + b
	fmt.Println("Suma : ", result)
	//resta
	result = a - b
	fmt.Println("Resta : ", result)
	//multiplicacion
	result = a * b
	fmt.Println("Multiplicacion : ", result)
	//division
	var div float64 = 3.0 / 2.0
	fmt.Println("Division : ", div)

	//modulo

	result = 3 % 2
	fmt.Println("Modulo", result)
}
