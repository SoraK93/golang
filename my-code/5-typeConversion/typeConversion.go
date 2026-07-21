package typeconversion

import "fmt"
// Type Conversion: Convert a value to another type.
// type(value): name of the time. changes the value to given type name
// We cannot use different types of values together

func Conversion() {
	speed := 100
	force := 2.5

	speed = speed * int(force)

	fmt.Println(speed)
}