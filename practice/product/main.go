package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	// your code goes here
	switch product {
	case "apples" :
		return price * 0.9
	case "bananas" :
		return price * 0.8
	case "orange", "oranges" :
		return price
	default:
		return price
	}
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
