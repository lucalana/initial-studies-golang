package main

import (
	"errors"
	"fmt"
)

func main() {
	// Funções em go podem retornar mais de um valor

	var num1 int = 351
	var num2 int = 3
	resultado, isError := sum(num1, num2)
	if isError != nil {
		fmt.Println(isError)
	}
	fmt.Println(resultado)
}

func sum(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
