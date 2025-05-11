package main

import (
	"fmt"
)

func main() {
	// Funções em go podem retornar mais de um valor

	saudacao := func(nome string) {
		fmt.Println("Olá " + nome + "!")
	}

	saudacao("Luca")
	saudacao("José")
	saudacao("Ana")
}
