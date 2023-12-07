package main

import (
  "fmt"
  "strings"
)

func main() {
  names_array := [4]string{"Kurt","Janis","Jimi","Amy"}
  fmt.Println(names_array)

  fmt.Printf("names_array type %[1]T of value %[1]v",names_array)

  names_slice := []string{"foo","bar"}
  fmt.Println(names_slice)

  // will not show a type slice, instace shows like a []type
  fmt.Printf("names_slice type %[1]T of value %[1]v\n",names_slice)

  //i := 6
  //fmt.Println("Name in 6: ",names_array[6]) // will not compile
  //fmt.Println("Name in 6: ",names_slice[6]) // compile but panic at runtime
  //fmt.Println("Name in 6: ",names_slice[i]) // compile but panic at runtime
  
  // append a slice
  names_slice = append(names_slice,"Jonathan")
  fmt.Println(names_slice)

  // append multiple slices
  slice1 := []string{"foo1","bar1"}
  slice2 := []string{"foo2","bar2"}

  for _, name := range slice2 {
    slice1 = append(slice1,name)
  }

  fmt.Println("Slice1: ", slice1, " slice2", slice2)

  // another way of doing the above
  slice1 = []string{"foo1","bar1"}
  slice2 = []string{"foo2","bar2"}

  slice1 = append(slice1,slice2...)

  fmt.Println("Slice1: ", slice1)

  // Slices have Lenght and Capacity
  fmt.Println("Slice1 capacity: ", cap(slice1))
  fmt.Println("Slice1 lenght: ", len(slice1))

  // slice capacity over a million iterations
  var s1 []int
  hat := cap(s1)
  for i := 0; i < 1_000_000; i++ {
    s1 = append(s1, i)
    c := cap(s1)
    if c != hat {
      fmt.Println(hat,c)
    }
    hat = c
  }


  // to change a subset of slace without changing the original slice
  // you need to copy
  names := [4]string{"Kurt","Janis","Jimi","Amy"}
  fmt.Println("Names: ",names)
  subset := make([]string,3)
  copy(subset, names[:3])
  fmt.Println("Subset: ", subset)
  for i,g := range subset {
    subset[i] = strings.ToUpper(g)
  }
  fmt.Println("New Subset: ",subset)
  fmt.Println("Names: ", names)
  
}
