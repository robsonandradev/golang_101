package basics

import (
  "fmt"
  "errors"
)

func RunErrorHandlers() {
  var nums []float32 = []float32{-2, 1, 4, 10, -3}
  for _, n := range nums {
    //r, err := errorNew(n)
    callNewError(n, errorNew)
    callNewError(n, appError.customError)
  }
  err := fmtErrorsf()
  fmt.Println(err)
}

func errorNew(num float32) (float32, interface{}) {
  if num < 0 {
    return 0, errors.New("Num less than zero can't be divide by 2!")
  }
  return num / 2, nil
}

//func callNewError(n float32, fn func(num float32) (float32, interface{})) {
func callNewError(n float32, fn appHandler) {
  //r, err := errorNew(n)
  r, err := fn(n)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(r)
}

func fmtErrorsf() error {
  return fmt.Errorf("number %g less than zero is invalid!", float32(-5))
}

type appError struct {
  error
  message string
  code    uint16
}

type appHandler func(num float32) (float32, interface{})

func (fn appError) customError(num float32) (float32, interface{}) {
  r, err := divByTwo(num)
  if err != nil {
    return r, &appError{err, "page not found", 404}
  }
  return r, nil
}

func divByTwo(num float32) (float32, error) {
  if num < 0 {
    return 0, fmt.Errorf("number %g less than zero is invalid!", num)
  }
  return num / 2, nil
}

//func (fn appError) Error(num float32) string {
  //return fmt.FormatString("number %g less than zero is invalid!", num)
//}
