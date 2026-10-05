package main

import "fmt"

func Variable() {
	/*
		Naming rules:
		1. If a name starts with an uppercase letter, it becomes exported like Speed
		2. name can start with an underscore letter like _speed
		3. name cannot start with a number(1Speed)/punctuation(!Speed)
		4. name cannot have punctuation in middle or end too like Spe!ed, Speed!
		5. cannot use reserved Go keywords as variable name
	*/

	// --- in Go: every variable and value has a type ---
	// variable type cannot be changed later. since, go is a strongly typed language.

	var speed int
	// var - lets declare a variable
	// speed - name/identifier of the variable
	// int - static type of the variable

	fmt.Println(speed)
}