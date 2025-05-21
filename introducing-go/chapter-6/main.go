package main

import "fmt"

func main() {
	// Exercise swap values
	l := 2
	p := 3
	fmt.Println(l, p)
	swap(&l, &p)
	fmt.Println(l, p)

	// Pointers
	// The * and & operator
	// Quando usamos o * estamos referenciando ao valor que esta armazenado naquele endereço de memória
	// Quando usamos o & estamos referenciando o endereço de memória daquela variável
	t := 5
	zero(t)
	fmt.Println(t) // t vai continuar sendo 5 por que o x definido dentro da função representa outro local da memória
	zeroP(&t)
	fmt.Println(t) // t vai ser alterado por que agora estamos passando o endereço na memória de t e alterando o valor que esta armazenado

	// new
	// Uma outra forma de criar um ponteiro
	nova := new(int)
	fmt.Println("Outra forma de conseguir ponteiro")
	fmt.Println(*nova)
	one(nova)
	fmt.Println(*nova)

	// -----------------------------------------------------------------------------------------------
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	x, y := f()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))
	// Uma forma de descompactar array
	fmt.Println(add(xs...))

	// Closure
	// Funções criadas assim tem acesso a variáveis de mesmo escopo que ela
	soma := func(x, y int) int { return x + y }
	fmt.Println(soma(2, 3))

	// Defer
	// Ele vai rodar o que estiver após defer quando a função principal terminar MESMO SE ACONTECER QUALQUER ERRO
	defer second()
	first()

	// Panic, Recover
	// Panic é para causar um erro de execução e o recover é para lidar com o erro de execução lançado pelo panic
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func one(x *int) {
	*x = 1
}

func zeroP(x *int) {
	*x = 0
}

func zero(x int) {
	x = 0
}

func first() {
	fmt.Println("Primeira")
}

func second() {
	fmt.Println("Segunda")
}

func average(xs []float64) float64 {
	total := 0.0
	for _, n := range xs {
		total += n
	}
	return total / float64(len(xs))
}

// O tipo do retorno pode ter nome
func f2() (r int) {
	r = 1
	return r
}

// Pode ser retornado mais de um valor
func f() (int, int) {
	return 5, 6
}

// Função que recebe vários parâmetros

func add(args ...float64) float64 {
	total := 0.0
	for _, n := range args {
		total += n
	}
	return total
}
