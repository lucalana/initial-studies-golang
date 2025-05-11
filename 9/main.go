package main

import (
	"fmt"
)

func main() {
	// Funções em go podem retornar mais de um valor

	var num1 int = 351
	var num2 int = 3
	var resultado int = sum(num1, num2, num2, num2)
	slice := []int{1, 2, 3, 4, 5}
	var resultado1 int = sum(slice...)

	fmt.Println(resultado)
	fmt.Println(resultado1)
}

func sum(numeros ...int) int {
	var valor int
	for _, n := range numeros {
		valor += n
	}
	return valor
}
