package main

import "fmt"

type Endereco struct {
	Cidade string
	Estado string
}

type Client struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c Client) toggleAtivar() {
	c.Ativo = !c.Ativo
}

func main() {
	pessoa := Client{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	pessoa.toggleAtivar()
	fmt.Println(pessoa.Ativo)

	fmt.Println(pessoa.Nome)
}
