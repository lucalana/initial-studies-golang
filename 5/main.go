package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(len(myArray))
	fmt.Println(myArray[2])

	for i, v := range myArray {
		fmt.Printf("Indice: %v ", i)
		fmt.Printf("Valor: %v\n", v)
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Indice: %v ", i)
		fmt.Printf("Valor: %v\n", myArray[i])
	}
}
