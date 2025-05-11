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

func main() {
	pessoa := Client{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	fmt.Println(pessoa.Nome)
}
