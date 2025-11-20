package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice) // Output: [1 2 3 4 5]
}
