package main

import (
	"fmt"
	"meuprojeto/matematica"
)

// Para que go identifique os moddulos que estamos criando devemos usar o comando go mod init nome _modulo
// E para que as varaiveis e metodos definifidos fiquei visiveis devemos usar letra maiusculo
// E para que eles não fiquem visiveis letra minuscula

func main() {
	s := matematica.Soma(2, 2)
	fmt.Println("Resultado: ", s)
}
