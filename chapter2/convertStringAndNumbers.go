package main

import (
  "fmt"
  "log"
  "strconv"
)


func main() {

  // converte a string to a negative integer
  i, err := strconv.ParseInt("-42",0,64)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("-42: ",i)


  // parsing a octal integer
  i, err = strconv.ParseInt("0x2A", 0, 64)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("0x2A: ",i)
  

  // parsing an unsigned integer
  u, err := strconv.ParseUint("42",0,64)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("42 (uint): ",u)

  // parsing a float
  f, err := strconv.ParseFloat("42.12345",64)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("42.12345: ",f)

  // converting a string to an integer
  a, err := strconv.Atoi("42")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%[1]v, [%[1]T]\n",a)

}
