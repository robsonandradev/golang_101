package main

import "fmt"

type iGeneric interface {}

func RunAssertation() {
  var i iGeneric = "string"
  assertationMatch(i)
  assertationUnMatch(i)
}

func assertationMatch(i iGeneric) {
  s, ok := i.(string)
  fmt.Println(s, ok)
}

func assertationUnMatch(i iGeneric) {
  s, ok := i.(float32)
  fmt.Println(s, ok)
}
