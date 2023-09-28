package main

import "fmt"

func main() {
  l := LinkedList{}
  l.Append("1")
  l.Append("2")
  l.Append("3")
  l.Append("4")
  l.Display()
  l.DisplayReversed()
}

type Node struct {
  data  string
  prior *Node
  next  *Node
}

type LinkedList struct {
  head *Node
  last *Node
}

func (l *LinkedList) Append(data string) {
  currentNode := l.head
  if l.head == nil {
    l.head = &Node{data: data}
    return
  }
  for currentNode.next != nil {
    currentNode = currentNode.next
  }
  newNode := &Node{data: data, next: nil}
  currentNode.next = newNode
  newNode.prior = currentNode
  l.last = newNode
}

func (l *LinkedList) Display() {
  currentNode := l.head
  items := []string{}
  for currentNode != nil {
    items = append(items, currentNode.data)
    currentNode = currentNode.next
  }
  fmt.Println(items)
}

func (l *LinkedList) DisplayReversed() {
  currentNode := l.last
  items := []string{}
  for currentNode != nil {
    items = append(items, currentNode.data)
    currentNode = currentNode.prior
  }
  fmt.Println(items)
}
