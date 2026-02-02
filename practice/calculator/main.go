package main

import "fmt"

func calculate(a int, b int) []float64 {
  // your code goes here
  v1 := float64(a)
  v2 := float64(b)
	return [] float64{v1+v2, v1-v2, v1*v2, v1/v2}
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}