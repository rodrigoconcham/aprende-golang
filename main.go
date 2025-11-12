package main

import "fmt"

type Celsio float64

func (c Celsio) ToFahrenheit() float64 {
    return float64(c * 9/5 * 32)
}

func main() {
    temp := Celsio(25)
    
    fmt.Println(temp.ToFahrenheit())
}