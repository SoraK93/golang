package main

import "fmt"

func ConvertType() {
	/*
		Type Conversion: converts a value to another type
		Syntax: typeName(value)
		typeName = changes the type of a given value into this type
		value = value to be converted

		Conversion can change value itself not just its type
	*/

	speed := 100 // speed is int
	force := 2.5 // force is float64

	// speed = speed * force // gives "Compiler Error: mismatch types"
	speed = speed * int(force) // 2.5 -> 2
	fmt.Println(speed)
	fmt.Println(force, int(force))

	// Order of conversion is very important
	speed = int(float64(speed) * force) // int(100.0 * 2.5) => int(250.0) => 250
	fmt.Println(speed)
}