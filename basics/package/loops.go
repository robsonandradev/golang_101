package basics

import "fmt"

func RunLoops() {
  whileLoop()
  classicFor()
  infinitLoop()
  continueLoop()
  loopWithRange()
}

func whileLoop() {
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
}

func classicFor() {
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
}

func infinitLoop() {
    for {
        fmt.Println("loop")
        break
    }
}

func continueLoop() {
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

func loopWithRange() {
  arr := [7]uint8{0, 1, 2, 3, 4, 5, 6}
  for i, v := range arr {
    fmt.Println("index: ", i, "value: ", v)
  }
}
