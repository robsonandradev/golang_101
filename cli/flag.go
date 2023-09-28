package main

import (
  "fmt"
  "flag"
)

func main() {
  flag.Parse()
  args := flag.Args()
  fmt.Println(args)
  for _, v := range args {
    fmt.Println(v)
  }
}
