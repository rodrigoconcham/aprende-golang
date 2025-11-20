package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println(slice) // Output: [20 30 40]
}
