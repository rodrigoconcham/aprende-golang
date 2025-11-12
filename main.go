package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            break
        }
        fmt.Println(i) // Output: 0, 1, 2
    }
}