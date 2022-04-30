package main

import "fmt"

func main() {
	varFloat := 5.8
	varInt1 := 1
	varInt2 := 2
	varInt3 := 3

	var result float64

	result = varFloat - float64(varInt1)
	fmt.Print("La operación ", varFloat, " - ", varInt1, " es igual a ", result)

	result *= float64(varInt2)
	fmt.Print(" x ", varInt2, " es igual a ", result)

	result /= float64(varInt3)
	fmt.Println(" / ", varInt3, " es igual a ", result)
}
