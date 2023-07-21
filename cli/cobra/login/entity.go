package login

type ILogin interface {
  doLogin(email string, password string) (Login, error)
}

type Login struct {
  email    string
  password string
}
