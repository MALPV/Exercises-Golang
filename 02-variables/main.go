package main

import "fmt"

func main() {

	var nombre string
	nombre = "Marcos"

	var apellido string = "Palma Valenzuela"

	edad := 24

	fmt.Println(nombre, apellido, "edad:", edad)

	var nroEntero int      // Valor por defecto 0
	var nroDecimal float64 // Valor por defecto 0
	var cadenaTexto string // Valor por defecto ""
	var condicional bool   // Valor por defecto false

	fmt.Println(nroEntero, nroDecimal, cadenaTexto, condicional)

	const pi = 3.141592
	fmt.Println(pi)
}
