package main

import "fmt"

func runPointer() {
  declarePointer()
}

func declarePointer() {
  var pointer *int
  age := 2
  pointer = &age // & used to pass the memory address
  fmt.Println(pointer)
  fmt.Println(*pointer) // * used to get the value storaged in this memory address
}
