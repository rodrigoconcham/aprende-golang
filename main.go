   package main


import "fmt"

func decirAdios() string {
    return "Adios"
}



func main() {

	despedida :=decirAdios()
   fmt.Println(despedida)

  divided := divide(1,2)
  fmt.Println(divided)
   



	 
}

func divide(a , b int ) int {

  if b==0  {
    return 0
  }

  return a / b 
}
