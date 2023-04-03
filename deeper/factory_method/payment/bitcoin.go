package payment

import "fmt"

type BitCoin struct {
  WalletPublicKey string
}

func (b BitCoin) Pay(val float32) (float32, error) {
  if b.WalletPublicKey == "" {
    return 0, fmt.Errorf("You must set WalletPublicKey")
  }
  fmt.Printf("send %f to %s wallet \n", val, b.WalletPublicKey)
  fmt.Println("processing payment")
  fmt.Println("payment confirmed")
  return 0, nil
}
