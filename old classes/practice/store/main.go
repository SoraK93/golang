package main

import "fmt"

// Declare the Expense struct here
// your code goes here
type Expense struct {
	name string
	amount float64
	date string
}

// Implement the Total method to calculate the total amount spent
// your code goes here
func Total(exp []Expense) float64 {
	sum := 0.00
	for _, value := range exp {
		sum += value.amount
	}
	return sum
}

// Implement the getName method on the Expense struct here
// your code goes here
func (exp Expense) getName() string {
	return exp.name
}

func main() {
	expenses := []Expense{
		{"Grocery", 50.0, "2022-01-01"},
		{"Gas", 30.0, "2022-01-02"},
		{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}