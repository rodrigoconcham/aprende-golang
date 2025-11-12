package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if j == 2 {
                continue
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
