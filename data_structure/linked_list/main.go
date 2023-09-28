package main

import "fmt"

func main() {
  l := LinkedList{}
  l.Append("1")
  l.Append("2")
  l.Append("3")
  l.Append("4")
  l.Display()
}

type Node struct {
  data string
  next *Node
}

type LinkedList struct {
  head *Node
}

func (l *LinkedList) Append(data string) {
  currentNode := l.head
  if l.head == nil {
    l.head = &Node{data: data, next: nil}
    return
  }
  for currentNode.next != nil {
    currentNode = currentNode.next
  }
  currentNode.next = &Node{data: data, next: nil}
}

func (l *LinkedList) Display() {
  currentNode := l.head
  items := []string{}
  for currentNode != nil {
    items = append(items, currentNode.data)
    currentNode = currentNode.next
  }
  fmt.Print(items)
}
