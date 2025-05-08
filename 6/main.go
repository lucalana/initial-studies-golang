package main

import "fmt"

func main() {
	// Seria como um array dinâmico
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Tamanho=%v Capacidade=%v Valores=%v\n", len(slice), cap(slice), slice)

	fmt.Printf("Tamanho=%v Capacidade=%v Valores=%v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("Tamanho=%v Capacidade=%v Valores=%v\n", len(slice[:1]), cap(slice[:1]), slice[:1])

	fmt.Printf("Tamanho=%v Capacidade=%v Valores=%v\n", len(slice[1:3]), cap(slice[1:3]), slice[1:3])

	// Ao adicionar um elemento no slice e o tamanho suportado ja passou ele vai dobrar de tamanho
	slice = append(slice, 12)
	slice = append(slice, 13)

	fmt.Printf("Tamanho=%v Capacidade=%v Valores=%v\n", len(slice), cap(slice), slice)
}
