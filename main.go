package main

import (
    "fmt"
    "time"
)

func Saluda() {
    fmt.Println("Hola desde un goroutine")
}

func main() {
    go Saluda()
    time.Sleep(time.Second)
    
    fmt.Println("main function completed")
}


