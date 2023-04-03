package main

import (
  "bufio"
  "fmt"
  "strconv"
  "strings"
  "os"
  "robsonandradev/factory_method/payment"
)

func main() {
  option := readStdin()
  pm, err := payment.ChoosePaymentMethod(option)
  if err != nil {
    panic(err)
  }
  pm.Pay(100.0)
}

func readStdin() (payment.PaymentMethod) {
  r := bufio.NewReader(os.Stdin)
  fmt.Println("1 - bitcoin")
  fmt.Println("2 - credit card")
  fmt.Fprint(os.Stderr, "choose your payment method: ")
  s, _ := r.ReadString('\n')
  optInt, _ := strconv.Atoi(strings.TrimSpace(s))
  return payment.PaymentMethod(optInt)
}
