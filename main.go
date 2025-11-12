package main

import "fmt"

func main() {
    letter := "a"
    switch letter {
    case "a", "e", "i", "o", "u":
        fmt.Println("Vocal")
    default:
        fmt.Println("Consonante")
    }
}