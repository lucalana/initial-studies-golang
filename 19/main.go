package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nomes := []string{"jose", "ana", "paulo", "flinstos"}

	// É possível esconder o indice colocando uma "_"
	for i, v := range nomes {
		fmt.Println(i)
		fmt.Println(v)
	}

}
