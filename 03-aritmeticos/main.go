package main

import "fmt"

func main() {

	nro1 := 7
	nro2 := 2

	result := nro1 + nro2
	fmt.Println("Suma:", result)

	result = nro1 - nro2
	fmt.Println("Resta:", result)

	result = nro1 * nro2
	fmt.Println("Multiplicación:", result)

	result = nro1 / nro2
	fmt.Println("División:", result)

	fmt.Println("División Exacta:", divisionExacta(nro1, nro2))

	result = nro1 % nro2
	fmt.Println("Modulo:", result)
}

func divisionExacta(nro1 int, nro2 int) float64 {
	var div float64 = float64(nro1) / float64(nro2)
	return div
}
