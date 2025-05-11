package main

import "fmt"

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func (c *Cliente) togglePlano() {
	c.ativo = !c.ativo
}

func main() {

	cliente1 := Cliente{
		nome:  "Luca",
		idade: 45,
		ativo: true,
	}

	fmt.Printf("Nome %v Ativo %v \n", cliente1.nome, cliente1.ativo)
	cliente1.togglePlano()
	fmt.Printf("Nome %v Ativo %v", cliente1.nome, cliente1.ativo)
}
