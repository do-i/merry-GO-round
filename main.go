package main

import "fmt"
import math "github.com/do-i/merry-go-round/math"

func main() {
  fmt.Println("this is main func in factorial pkg")
  fmt.Printf("Factorial of %d = %d\n", 5, math.FactorialOf(5))
  fmt.Printf("Combination of C(%d, %d) = %d\n", 5, 3, math.CombinationOf(5, 3))
}