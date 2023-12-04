package main

import (
  "fmt"
)

type User struct {
  Name string
  Age int
}

func main() {
  fmt.Printf("%+d\n", -123)
  fmt.Printf("%+d\n", 123)


  // Padding numbers

  var d = 5987

  fmt.Printf("Padding: %7d\n",d)
  fmt.Printf("Padding: %3d\n",d)
  fmt.Printf("Padding: %-10d\n",d)

  u := User{ Name: "Julio", Age: 35 }

  fmt.Printf("Type of u: %T\n",u)
  fmt.Printf("Value of u: %+v\n",u)
  fmt.Printf("Value of u in Go-syntax format: %#v\n",u)


  fmt.Printf("Value of u: %+[1]v and type %[1]T",u)
}
