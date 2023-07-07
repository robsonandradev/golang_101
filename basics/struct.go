package main

import "fmt"

type Employee struct {
  name   string
  age    uint8
  salary float32
}

func runStructSample() {
  //initializeWithoutNameFields()
  //initializeWithNamedFields()
  //initializeWithSomeFieldsNamed()
  pointerWithStructs()
}

func deleteItems(employee *Employee) {
  employee.name = "john Travolta"
  fmt.Println(employee)
}

func pointerWithStructs() {
  employee := Employee{"john wick", 35, 100000.00}
  fmt.Println(employee)
  deleteItems(&employee)
  fmt.Println(employee)
}

func initializeWithoutNameFields() {
  employee := Employee{"john wick", 35, 100000.00}
  fmt.Println(employee)
}

func initializeWithNamedFields() {
  employee := Employee{
    name: "john wick",
    age:  35,
    salary: 100000.00,
  }
  fmt.Println(employee)
}

func initializeWithSomeFieldsNamed() {
  employee := Employee{name: "john wick", age: 35}
  fmt.Println(employee)
}
