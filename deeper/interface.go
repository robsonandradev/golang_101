package main

import "fmt"

type iPayment interface {
  Pay(value float32) (float32, error)
}

type CreditCard struct {
  owner string
  cardNumber string
}

func (card CreditCard) Pay(value float32) (float32, error) {
  fmt.Println("Processing payment")
  fmt.Println("checking card limit")
  _, err := fmt.Printf("call payment gateway api - value sent: %f \n", value)
  if err != nil {
    return 0, err
  }
  fmt.Println("success paid")
  return 0, nil
}

func Payment() {
  change, err := execPayment(CreditCard{"jonh wick", "55123"})
  if err != nil {
    panic(err)
  }
  fmt.Printf("change is %f", change)
}

func execPayment(payment iPayment) (float32, error) {
  return payment.Pay(100)
}
