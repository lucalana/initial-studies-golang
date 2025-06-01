package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		number := i + 1
		fmt.Println(number)
		if number%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}
	i := 2
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown number")
	}
}
