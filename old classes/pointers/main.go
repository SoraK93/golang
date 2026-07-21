package main

import (
	"fmt"
	"strings"
)

func main() {
	// scope
	y := 20
	modify(y)
	fmt.Println(y)

	// dereference and assign new value
	derefAssign()
	derefAssign2()

	// dereference and do calculation
	calOnDeref()
	calOnDeref2()

	// passing by reference
	passByRef()
}

func modify(y int) {
	y += 15
	fmt.Println(y)
}

func derefAssign() {
	x := [3]int{10, 20, 30}
	fmt.Printf("%v \n", x)
	(*&x)[0] = 100
	fmt.Printf("%v \n", x)
}

func derefAssign2() {
	s := "hello"
	var ptr *string = &s
	fmt.Println(s)
	*ptr += strings.ToUpper(s)
	fmt.Println(s)
}

func calOnDeref() {
	var y int
	var ptr *int = &y
	*ptr = 0
	fmt.Println(y)
	*ptr += 5
	fmt.Println(y)
}

func calOnDeref2() {
	s := 100
	var ptr *int = &s
	fmt.Println(s)
	*ptr += 100
	fmt.Println(s)
}

func passByRef() {
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modifySlice(arr...)
	fmt.Println(arr)

	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	fmt.Println(ascii_codes)
	modifyMap(ascii_codes)
	fmt.Println(ascii_codes)
}
func modifySlice(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}
func modifyMap(s map[string]int) {
	s["A"] = 100
}
