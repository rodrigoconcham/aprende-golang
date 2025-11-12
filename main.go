package main

import "fmt"

func main() {
    defer fmt.Println("Esto esta impreso ultimo")
    fmt.Println("Esto esta impreso primero")
}
