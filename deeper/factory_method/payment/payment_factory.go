package payment

import "fmt"

type PaymentMethod uint8

const (
  BITCOIN_METHOD     PaymentMethod = 1
  CREDIT_CARD_METHOD PaymentMethod = 2
)

func ChoosePaymentMethod(option PaymentMethod) (IPayment, error) {
  switch option {
  case BITCOIN_METHOD:
    return bitcoinMethod(), nil
  case CREDIT_CARD_METHOD:
    return creditCardMethod(), nil
  }
  return nil, fmt.Errorf("Unknow payemnt method")
}

func creditCardMethod() IPayment {
  creditcard := CreditCard{Owner: "john wick", CardNumber: "552512345"}
  return creditcard
}

func bitcoinMethod() IPayment {
  publicKey := "Mzc3MGI4MTQ5ODFlN2I3ZTk4MzdjOGRiNmNhODQwMGMwYzcwNjI3ZCAgLQo="
  bitcoin := BitCoin{WalletPublicKey: publicKey}
  return bitcoin
}
