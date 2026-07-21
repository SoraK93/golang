package main

import "fmt"

func main() {
	var cast, name string

	fmt.Print("Please enter caster name and magic name: ")
	count, err := fmt.Scanf("%s %s", &name, &cast)
	fmt.Printf("%v\n", count)
	fmt.Printf("%v\n", err)
	fmt.Printf("Magic user %s casts %s", name, cast)
}