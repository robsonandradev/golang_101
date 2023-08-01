package main

func main() {
  book1 := BookFactory("Paulo Coelho", "O demonio e sr green", "o cabrunco perturba!")
  replaced := book1.ReplaceWordInText("cabrunco", "diabo")
  bPrinter := BookPrinterFactory()
  bPrinter.ShowText(replaced)
  bPrinter.ShowText(book1.GetText())
  //isCabrunco = book1.IsWordInText("diabo")
  //fmt.Println(isCabrunco)
}
