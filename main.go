package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["apple"] = 2
	m["banana"] = 3
	fmt.Println(m["apple"]) // Output: 2
}
