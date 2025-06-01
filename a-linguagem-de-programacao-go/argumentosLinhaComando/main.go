package main

import (
	"fmt"
	"os"
	"strings"
)

// Exibe argumentos de linha de comando
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// Ou maneira de realizar loop for
	for indice, arg := range os.Args {
		fmt.Println(indice, arg)
	}

	//Caso args tenha muitos argumentos vai ser custoso para o programa juntar tudo em apenas uma variável então podemos usar a funçãoo join
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println(s)
}
