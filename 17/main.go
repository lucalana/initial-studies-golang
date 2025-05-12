package main

func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"jose": 1000, "flinston": 2000}
	a := map[string]float64{"jose": 123.33, "flinston": 234.12}

	println(soma(m))
	println(soma(a))

}
