package main

import (
  "fmt"
)


func main() {


  // Go only have FOR for loops
  // classic
  for i := 1; i <= 10; i++ {
    fmt.Println("Number: ",i)
  }

  
  // While with FOR
  j := 0
  for {
    j++

    fmt.Println("Number: ",j)
    
    if ( j > 10 ) {
      break
    }
  }

  // Using range
  names := []string{"macos","debian","almalinux","iOS","Android"}
  fmt.Println("Names:")
  for idx,name := range names {
    fmt.Printf("\t%[1]v %[2]v %[2]T\n",idx,name)
  }

}
