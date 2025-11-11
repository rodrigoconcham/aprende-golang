package main

import (
	"errors"
	"fmt"
)

func decirHola() {
	fmt.Println("Hola!")
}

func main() {

	decirHola()

	suma := add(2, 3)

	fmt.Println(suma)

	division, err := dividir(1, 1)

	if err != nil {
		fmt.Println("hay error")
	}

	fmt.Println(division)

	if y:=10; y > 5 {
        fmt.Println("y es mayor de 5")

	}
}

func add(x, y int) int {
	return x + y
}

func dividir(x, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("no puede ser zero")
	}

	return x / y, nil

}
