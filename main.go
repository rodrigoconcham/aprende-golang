package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4}
    for i, num := range nums {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
}
