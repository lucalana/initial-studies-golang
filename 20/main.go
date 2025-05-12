package main

import "fmt"

func main() {
	x := 4

	if x > 10 {
		println("maior")
	} else if x == 10 {
		println("igual")
	} else {
		println("menor")
	}

	switch x {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	default:
		fmt.Println("Outro número")
	}

	switch {
	case x > 10:
		fmt.Println("Maior que 10")
	case x == 10:
		fmt.Println("Igual a 10")
	default:
		fmt.Println("Menor que 10")
	}

}
