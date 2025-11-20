package main

import "fmt"

func main() {
	var arr [5]int // [0,0,0,0,0]
	arr[0] = 10    // [10,0,0,0,0]
	arr[1] = 20    // [10,20,0,0,0]

	fmt.Println(arr)    // Output: [10 20 0 0 0]
	fmt.Println(arr[1]) // Output: 20
	fmt.Println(arr[2])
}
