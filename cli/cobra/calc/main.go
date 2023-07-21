package calc

import "fmt"

type Calc struct {}

func (c *Calc) Sum(n1 int, n2 int) int {
  return n1 + n2
}

func (c *Calc) Subtraction(n1 int, n2 int) int {
  return n1 - n2
}

func (c *Calc) Multiply(n1 int, n2 int) int {
  return n1 * n2
}

func (c *Calc) Division(n1 float32, n2 float32) (float32, error) {
  if n2 <= 0 {
    return 0, fmt.Errorf("Couldnt div by 0 or negative numbers %f", n2)
  }
  return n1 / n2, nil
}

