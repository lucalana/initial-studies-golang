package main

import "fmt"

var y string = "Hello, word"
var (
	a = 5
	b = 1
	c = 6
)

// Não pode ser alterado
const PI float64 = 3.14345

func main() {
	var x string
	x = "Hello, word"
	name := "luca"

	fmt.Println(name)
	fmt.Println(x)
	fmt.Println(y)
	f()

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println(input * 2)
}

func f() {
	fmt.Println(y)
}
