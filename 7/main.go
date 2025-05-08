package main

import "fmt"

func main() {
	// Maps parte de chave e valor no go
	// Não é garantido a order
	// idades := make(map[string]int)
	// marcas := map[string]string{}
	idades := make(map[string]int)
	idades["luca"] = 21
	idades["jose"] = 22
	idades["joao"] = 23

	marcas := map[string]string{"ranger": "Ford", "uno": "Fiat", "sandero": "renault"}

	for nome, idade := range idades {
		fmt.Printf("Nome: %v Idade: %v \n", nome, idade)
	}
	// Caso n queira usar o index declare ele como _
	for _, idade := range idades {
		fmt.Printf("Idade: %v \n", idade)
	}

	fmt.Println(marcas)
	fmt.Println(idades)
	fmt.Println(idades["luca"])
}
