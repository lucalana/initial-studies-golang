package main

import (
	"fmt"
	"math"
)

// Structs
// É um tipo que contem campos nomeados
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Interface
type Shape interface {
	area() float64
}

func totalArea(s ...Shape) float64 {
	var total float64 = 0
	for _, v := range s {
		total += v.area()
	}
	return total
}

func main() {
	// Formas de inicializar
	// var c Circle
	// c := new(Circle)
	// c := Circle{0, 0, 5}
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	// Se quiser um ponteiro para estrutura usamos o &
	// c := &Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r)

	fmt.Println(totalArea(&c, &r))

	fmt.Println(r.area())
	fmt.Println(c.area())

	p := Person{"Luca"}
	p.Talk()

	a := Android{p, "G3"}
	a.Talk()
}

// Tipos embutidos
type Person struct {
	Name string
}

func (p Person) Talk() {
	fmt.Println("Ola meu nome é", p.Name)
}

// Dessa forma o android pode acessar diretamente os metodos de Person ou então chamar Person e acessar os metodos
type Android struct {
	Person
	Model string
}

// Adicionando MÉTODO a uma estrutura
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
