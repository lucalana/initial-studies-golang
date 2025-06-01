package main

import "fmt"

func main() {
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)
	var x [5]int
	x[0] = 100
	x[1] = 100
	x[2] = 100
	x[3] = 100
	x[4] = 100
	var total float64 = 0

	// for i := 0; i < len(x); i++ {
	// 	total += float64(x[i])
	// }
	for _, v := range x {
		total += float64(v)
	}

	fmt.Println(total / float64(len(x)))
	fmt.Println(x)

	//Slice
	// Duas formas de criar slice
	var a []float64
	b := make([]float64, 5)

	// Adiciona elementos ao slice(Ele cria um slice novo a partir de um que ja existe e adiciona elementos nele)
	a = append(a, 2, 3, 4, 5)
	b = append(b, 3, 4)

	fmt.Println(a)
	fmt.Println(b)
	// Copia os elementos de um slice para outro
	// Obs: Ele sempre vai usar o tamanho do menor
	copy(a, b)
	fmt.Println(a)

	//Maps (hash-table)
	m := make(map[string]int)
	shortMap := map[string]int{
		"jose":     1,
		"alfredin": 34,
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	fmt.Println(elements)
	fmt.Println(shortMap)
	m["key"] = 10
	m["luca"] = 256

	fmt.Println(m)
	fmt.Println(m["key"])
	// Deletar itens do map
	delete(m, "key")
	fmt.Println(m)
	fmt.Println(m["key"]) // Retorna 0 quando não encontra
	// Mas existe uma outra forma de fazer esse verificação
	if name, ok := m["key"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println(ok)
	}

	// Find the smallest number in the slice
	numbers := []int{435, 34, 554, 64, 45, 46, 266, 235, 236, 1, 2351, 123, 564, 75, 7658, 688, 3}

	var small int
	for i, n := range numbers {
		if i == 0 {
			small = n
		}
		if n < small {
			small = n
		}
	}
	fmt.Println(small)
}
