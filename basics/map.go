package main

import "fmt"

func runMaps() {
  declareMap()
  initMapWithMake()
  initMapWithLieteral()
  addToMap()
  getFromMap()
  deleteFromMap()
}

func declareMap() {
  var employee map[string]string
  fmt.Println(employee)
}

func initMapWithMake() {
  employee := make(map[string]string)
  fmt.Println(employee)
}

func initMapWithLieteral() {
  employeeSalary := map[string]float32{
    "john": 1000.0,
    "mario": 1000.0,
  }
  fmt.Println(employeeSalary)
}

func addToMap() {
  employeeSalary := map[string]float32{
    "john": 1000.0,
    "mario": 1000.0,
  }
  employeeSalary["gerald"] = 1500.0
  fmt.Println(employeeSalary)
}

func getFromMap() {
  employeeSalary := map[string]float32{
    "john": 1000.0,
    "mario": 1000.0,
  }
  fmt.Println(employeeSalary["john"])
}

func deleteFromMap() {
  employeeSalary := map[string]float32{
    "john": 1000.0,
    "mario": 1000.0,
  }
  delete(employeeSalary, "john")
  fmt.Println(employeeSalary)
}
