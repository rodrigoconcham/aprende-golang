package main

import (
	"fmt"
)

func decirHola() {
	fmt.Println("Hola!")
}

func main() {

	decirHola()

	suma := add(2, 3)

	fmt.Println(suma)
}

func add(x, y int) int {
	return x + y
}
