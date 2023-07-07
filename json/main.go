package main

import (
  "bytes"
  "encoding/json"
  "fmt"
)

type Parent struct {
  ParentName string
  ParentAge uint
  Sex byte
}

type FamilyMember struct {
  Name string
  Age uint
  Parents [2]Parent
}

func buildEnzo() FamilyMember {
  father := Parent{ParentName: "joao", ParentAge: 55, Sex: 1}
  mather := Parent{ParentName: "maria", ParentAge: 25, Sex: 0}
  parents := [2]Parent{father, mather}
  enzo   := FamilyMember{Name: "enzo", Age: 5, Parents: parents}
  return enzo
}

func bytesToStr(b []byte) string {
  var buf bytes.Buffer
  buf.Write(b)
  return buf.String()
}

func structExample() {
  //b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
  enzo := buildEnzo()
  e, err := json.MarshalIndent(enzo, "", " ")
  if err != nil { panic(err) }
  eNoIndent, err := json.Marshal(enzo)
  if err != nil { panic(err) }
  eStr := bytesToStr(eNoIndent)
  fmt.Println(string(e[:]))
  fmt.Println(eStr)
}

func mapAsJson() {
  f := map[string]interface{}{
    "Name": "Enzo",
    "Age": 5,
    "Parents": []interface{}{
      "Joao",
      "Maria",
    },
  }
  fmt.Println(f)
}

func main() {
  // structExample()
  mapAsJson()
}
