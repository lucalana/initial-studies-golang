package main

import "fmt"

type ID int

var idade ID = 2

func main() {
	fmt.Printf("O tipo da idade é %T\n", idade)
	fmt.Printf("O valor da idade é %v\n", idade)
}
