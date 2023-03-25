package main

import "fmt"

func runVariables() {
  var age uint8 = 1 // uint8 2^8
  var name string = "john wick"
  var isMan bool = true

  var country string
  country = "abc"

  fmt.Println(name, age, isMan, country)
}
