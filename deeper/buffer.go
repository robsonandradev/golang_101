package main

import (
  "bytes"
  "fmt"
)

func RunBuffer() {
  var strBuffer bytes.Buffer
  fmt.Fprintf(&strBuffer, "It is a number %d, this is a string %v\n", 10, "Bridge")
  strBuffer.WriteString("[DONE]")
  fmt.Println("The string in the buffer is:", strBuffer.String())
}
