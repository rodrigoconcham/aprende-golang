package main

import  (
      "fmt"
  )


func main() {
    m := map[string]int{"manzana": 2, "banana": 3}
    fmt.Println(m["manzana"]) // Output: 2
  
    for key, value := range m {
        fmt.Printf("key es: %s; value es: %d", key, value)
    }

}