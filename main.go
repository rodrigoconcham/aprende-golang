package main

import "fmt"

func main() {
    number := 0
    switch number {
    case 0:
          fallthrough
    case 1:
        fmt.Println("Number is either 0 or 1")
    default:
        fmt.Println("Number is something else")
    }
}