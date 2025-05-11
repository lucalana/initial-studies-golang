package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10

	// Endereço da memória de a  onde o valor 10 esta guardado
	var ponteiro *int = &a

	// Alterando o valor desse endereço de memória pra 20
	*ponteiro = 20

	println(a)
	println(&a)
}
