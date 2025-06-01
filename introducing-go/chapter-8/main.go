package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Teste", "es"))
	fmt.Println(strings.Count("Teste", "e"))
	fmt.Println(strings.HasPrefix("Teste", "Te"))
	fmt.Println(strings.HasSuffix("Teste", "te"))
	fmt.Println(strings.Index("Teste", "e")) // Retorna o inddex da primeira
	fmt.Println(strings.Join([]string{"Jose", "do", "Morro"}, " "))
	fmt.Println(strings.Repeat("Jose", 5))
	// Ele recebe a string o que vai ser trocado pelo que vai ser trocado e quantidade de vezes
	// Se passar -1 ele vai trocar o a em todas as ocorrências encontradas
	fmt.Println(strings.Replace("aaaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-s-s-sd-sd-sd-a", "-"))
	fmt.Println(strings.ToUpper("tudo vai estar maiúsculo"))

	//Converter string para binário e vice versa
	arr := []byte("test")
	str := string([]byte("test"))
	fmt.Println(arr)
	fmt.Println(str)

	// Input / Output
}
