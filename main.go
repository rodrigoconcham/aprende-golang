package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	nombre := p.Name
	fmt.Println(nombre) // "Alice"
	fmt.Println(p)      // Output: {Alice 30}
}
