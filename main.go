package main

import "fmt"

func imprimirCualquierCosa(v interface{}) {
	fmt.Println(v)
}

func main() {
	imprimirCualquierCosa(42)
	imprimirCualquierCosa("hola")
	imprimirCualquierCosa(3.14)
}
