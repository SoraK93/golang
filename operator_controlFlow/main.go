package main

import "fmt"

func main() {
	comparison()

}

func comparison() {
	var a int = 12
	var b int = 12
	fmt.Println(a <= b)

	a = 20
	fmt.Println(a <= b)

	b = 100
	fmt.Println(a <= b)

	c := 0
	fmt.Println(a <= c)
}

func arithematic() {
	var (
		a float64 = 5.9
		b, c, d int = 2, 10, 4
	)

	fmt.Println((int(a) + b) * c / d)
}
