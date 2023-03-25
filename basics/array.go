package main

import "fmt"

func runArrays(){
  var arr [3]string
  arr = [3]string{"a", "b", "c"}
  fmt.Println(arr)
  arr2 := arrayAreValues(arr)
  fmt.Println(arr) // keeps the old value
  fmt.Println(arr2) // new array
}

func arrayAreValues(arr [3]string) [3]string {
  for i := 0; i < len(arr); i++ {
    arr[i] = fmt.Sprint(arr[i], "-", i)
  }
  return arr
}

func createSlice() {
  arr := createSliceWithMake()
  createSliceDirect()
  arr = appendSlice(arr)
  interateOverSlice(arr)
}

func createSliceWithMake() []string {
  arr := make([]string, 2, 3)
  fmt.Println(arr)
  return arr
}

func createSliceDirect() {
  arr := []string{"sonic", "john", "mario"}
  fmt.Println(arr)
}

func appendSlice(arr []string) []string {
  arr = append(arr, "a")
  arr = append(arr, "b")
  arr = append(arr, "c")
  fmt.Println(arr)
  return arr
}

func interateOverSlice(arr []string) {
  for idx, val := range arr {
    fmt.Println("index: ", idx, "value: ", val)
  }
}
