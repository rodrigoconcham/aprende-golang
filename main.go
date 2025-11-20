package main

import "fmt"

func main() {
	fmt.Println("Antes de goto")
	goto End
	fmt.Println("Esto no se ejecutara")
End:
	fmt.Println("Despues goto")
}
