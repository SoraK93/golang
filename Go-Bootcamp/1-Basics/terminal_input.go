package main

import ("fmt"; "os")

/*
Package 'os'
- allows you to access operating system functionalities

var Args []string
- Go puts the command-line arguments into this varaiable automatically

[]string - Args's type is a "slice of strings"

Example:
go run main.go hi yo
Here 'hi' & 'yo' becomes the arguments and ' ' is used as seperator to identify both as seperate strings. Both are then stored inside Args variable.
Args[0] = stores the temporary path to the currently executing program
Args[1] = 'hi'
Args[2] = 'yo'
*/

/*
Slices
- A slice can store multiple values
- each value inside slice is an unnamed variable
- We can use index expression to acccess each values inside a slice like Args[0].
Syntax = []dataType
*/

func TerminalInput() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3nd argument:", os.Args[3])

	fmt.Println("Number of items inside os.Args:", len(os.Args))

	// We can use go build -o <fileName> to create a build file, which we can call using ./<fileName>
}
