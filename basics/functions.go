package main

import "fmt"

func runFunctions() {
  variadicSum(1, 2, 3, 4)
  nums := []int8{1, 2, 3}
  variadicSum(nums...)
  anonymousFunc()
}

func variadicSum(nums ...int8) {
  fmt.Print(nums, " - Sum is: ")
  var total int8 = 0
  for _, n := range nums {
    total += n
  }
  fmt.Println(total)
}

func anonymousFunc() {
  fmt.Println(func() string {
    return "anon has retorn"
  }())
}
