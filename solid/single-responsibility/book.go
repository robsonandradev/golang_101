package main

import "strings"

type IBook interface {
  ReplaceWordInText(word string, replacementWord string) string
  IsWordInText(word string) bool
  GetText() string
}

type Book struct {
  author string
  title  string
  text   string
}

func BookFactory(author, title, text string) IBook {
  return &Book{author: author, title: title, text: text}
}

func (b *Book) ReplaceWordInText(word string, replacementWord string) string {
  b.text = strings.ReplaceAll(b.text, word, replacementWord)
  return b.text
}

func (b *Book) IsWordInText(word string) bool {
  return strings.Contains(b.text, word)
}

func (b *Book) GetText() string {return b.text}
