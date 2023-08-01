package main

import "fmt"

type IBookPrinter interface {
  ShowText(t string)
}

type BookPrinter struct {}

func BookPrinterFactory() IBookPrinter {
  return &BookPrinter{}
}

func (b *BookPrinter) ShowText(t string) {
  fmt.Println(t)
}
