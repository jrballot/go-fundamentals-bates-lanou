package main

import (
  "fmt"
)

func main() {
  // _ is a placeholder that will ignore the returned value
  _, f,b,s := Values()

  fmt.Println(f,b,s)


  const gopher string = "gopher" 
  fmt.Printf("Constant gopher type %T and value %s\n",gopher,gopher)
  //gopher = "can't overwrite a constant, this will fail" 

}


func Values() (int, float64, bool, string) {
  return 42, 3.14, true, "hello, I'm a string" 
}
