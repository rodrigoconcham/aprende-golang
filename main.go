package main

import "fmt"

func main() {
	x := 10
	p := &x
	fmt.Println(p)  // Salida: Dirección de memoria de x (ej: 0xc000014088)
	fmt.Println(*p) // Salida: 10
}
