package main

import (
    "fmt"
    "time"
)


func printNumbers(n int) {
    for i := 1; i <= n; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers(5)
    go printNumbers(5)
    time.Sleep(time.Second) // Allow time for both goroutines
}
