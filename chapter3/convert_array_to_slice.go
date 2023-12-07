package main

import (
  "fmt"
)

func slicesOnly(names []string){
  for _, name := range names {
    fmt.Println(name)
  }
}

func main() {
  names := [4]string{"tony","niamh","alba","ian"}
  
  // will fail because names is a [4]string and not []string as expected by slicesOnly()
  //slicesOnly(names)
  // solution, convert that array into subset with [:] notation
  slicesOnly(names[:])
}
