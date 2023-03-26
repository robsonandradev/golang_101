package main

import "fmt"

func runConditionalStatements() {
  numbers := [3]int8{-1, 0, 1}
  for _, num := range numbers {
    msg := ifElseStatement(num)
    fmt.Println(msg)
  }
  for _, num := range numbers {
    msg := simpleSwitchCase(num)
    fmt.Println(msg)
  }
  for _, c := range [3]string{"!", "&", "$"} {
    if caseList(c) {
      fmt.Println("valid character")
    } else {
      fmt.Println("invalid character")
    }
  }
}

func ifElseStatement(num int8) string {
  if num > 0 {
    return "Positive number"
  } else if num < 0 {
    return "negative number"
  } else {
    return "got zero"
  }
}

func simpleSwitchCase(num int8) string {
  switch {
    case num > 0:
      return "Positive number"
    case num < 0:
      return "negative number"
  }
  return "got 0"
}

func caseList(character string) bool {
  switch character {
    case "", "%", "!", "#", "^":
      return false
  }
  return true
}
