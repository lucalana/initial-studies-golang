# Anotações sobre o que foi visto em cada aula

## Declaração e atribuição

Constantes uma vez declaradas não podem ter seu valor alterado. E uma variável quando declarada de um 
determinado tipo não pode ser alterado(fortemente tipado), além de ser possível declarar mais de uma variável de uma vez.

```go
    const myConst = "Valor que nunca muda"
	var myBool bool = true
    var (
		myName string = "Luca"
		myAge  int    = 25
	)
	myNumber := 25 // Variável definida com inferência
```

Quando declaramos variáveis sem passar algum valor elas vão receber alguns valores padrão

```go
	var myBool bool // Vai receber false
    var (
		myName string // Vai receber ""
		myAge  int    // Vai receber 0
	)
```

### Escopo

Variáveis definidas fora da função vão ter o escopo global para todo pacote, e variáveis definidas 
dentro da função só vão ser vistas dentro da própria função