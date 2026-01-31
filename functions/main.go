package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// normal function
	fmt.Println(returnCube(3))

	// function with multiple returns
	nums := [3]int{10, 20, 15}
	fmt.Println(calcSquare(nums[:]))

	// variadic function
	result := printStrings("Joe", "Monica", "Gunther")
	fmt.Println(result)

	// recersive function
	recursivePrint(5)
	fmt.Print("\n")

	// anonymous function
	anonymous()

	// higher-order function
	partial := partialSum(addHundred, 1, 2, 3, 4, 5)
	fmt.Println(partial)

	// defer statement
	deferEx()
}

func returnCube(n int) int {
	return n * n * n
}

func calcSquare(numbers []int) ([]int, bool) {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares, true

}

func printStrings(names ...string) (names_c []string) {
	names_c = []string{}
	for _, value := range names {
		names_c = append(names_c, strings.ToUpper(value))
	}
	return
}

func recursivePrint(n int) {
	if n == 0 {
		return
	}
	recursivePrint(n - 1)
	fmt.Print(n)
}

func anonymous() {
	x := func(s string) string {
		fmt.Println(strings.ToLower(s))
		return strings.ToLower(s)
	}("RacheL")
	fmt.Printf("%T \n", x)

	var (
		cube = func(i int) string {
			c := i * i * i
			return strconv.Itoa(c)
		}
	)
	n := cube(8)
	fmt.Printf("%T %v\n", n, n)
}

func addHundred(x int) {
	fmt.Println(x + 100)
}
func partialSum(add100 func(x int), x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	add100(sum)
	return 0
}

func printStr(str string) {
	fmt.Printf("%q ", str)
}
func printInt(i int) {
	fmt.Printf("%d ", i)
}
func printFloat(f float64) {
	fmt.Printf("%.2f ", f)
}
func deferEx() {
	printStr("browser")
	defer printInt(32)
	defer printFloat(0.24)
	printStr("chrome")
	printInt(90)
	defer printFloat(89)
	printInt(900)
}
