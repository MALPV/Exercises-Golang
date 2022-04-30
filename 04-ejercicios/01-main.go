package main

import "fmt"

func main() {

	var name string
	var age int
	var isAdult bool

	fmt.Println("Cual es tu nombre?")
	fmt.Scanf("%s\n", &name)

	fmt.Println("Cual es tu edad?")
	fmt.Scanf("%d\n", &age)

	if age > 18 {
		isAdult = true
	}

	fmt.Println("Bienvenido(a)", name)
	fmt.Println("Tu edad es:", age)

	if isAdult {
		fmt.Println("Ya eres un adulto(a)")
	}
}
