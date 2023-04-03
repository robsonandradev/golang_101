package payment

import "fmt"

type CreditCard struct {
  CardNumber string
  Owner      string
}

func (cc CreditCard) Pay(val float32) (float32, error) {
  if cc.Owner == "" {
    return 0, fmt.Errorf("You must set Card Owner")
  }
  if cc.CardNumber == "" {
    return 0, fmt.Errorf("You must set CardNumber")
  }
  fmt.Println("credit card chosen")
  fmt.Println("processing...")
  fmt.Println("Successful")
  return 0, nil
}
