package osargs

import (
	"fmt"
	"os"
)

func Greeter() {
	// Use build statement like "go build -o <file-name>" to build a named executable
	
	fmt.Printf("Path: %s\n", os.Args[0])
	fmt.Printf("1st Argument: %s\n", os.Args[1])
	fmt.Printf("2nd Argument: %s\n", os.Args[2])
	fmt.Printf("3rd Argument: %s\n", os.Args[3])

	fmt.Println("Number of items in os.Args:", len(os.Args))
}