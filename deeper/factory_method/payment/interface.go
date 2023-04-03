package payment

type IPayment interface {
  Pay(val float32) (float32, error)
}
