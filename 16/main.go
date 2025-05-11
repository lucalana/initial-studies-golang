package main

func main() {
	var y interface{} = "Hello Word!"

	res, ok := y.(int)

	if ok {
		println(res)
	} else {
		println("Deu ruim")

	}
}
