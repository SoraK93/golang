package main

import (
	"fmt"
	"strconv"
)

var count string = "1234"

func main() {
	newCount, err := strconv.Atoi(count)
	fmt.Println("Error: ", err)
	fmt.Printf("%d is of type %T\n", newCount, newCount)

	newString := strconv.Itoa(newCount)
	fmt.Printf("%s is of type %T\n", newString, newString)

	// When a int is converted, it just adds precision values
	newFloat := float64(newCount)
	fmt.Printf("%.2f is of type %T\n", newFloat, newFloat)

	// Number will lose the decimal value
	newInt := int(newFloat)
	fmt.Printf("%d is of type %T\n", newInt, newInt)
}
