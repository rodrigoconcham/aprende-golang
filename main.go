package main

import "fmt"

type Rectangulo struct {
    Width float64
    Height float64
}

func (r Rectangulo) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangulo{
        Width: 5,
        Height: 5,
    }
    
    fmt.Println(rect.Area()) // Output 25
}


