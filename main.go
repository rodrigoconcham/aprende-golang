package main

import "fmt"

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("Case 1")
        fallthrough
    case 2:
        fmt.Println("Case 2")
    }
}
