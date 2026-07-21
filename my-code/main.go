package main

import (
	"fmt"
	basics "soraK93/my-code-1/1-basics"
	pathseperator "soraK93/my-code-1/2-pathSeperator"
	typeconversion "soraK93/my-code-1/5-typeConversion"
	osargs "soraK93/my-code-1/7-OsArgs"
	"soraK93/my-code-1/demo1"
)

func main() {
	fmt.Println(demo1.DemoCode())
	fmt.Println(basics.FullVariables())
	fmt.Println(basics.ShortVariables())
	fmt.Println(pathseperator.Path())

	typeconversion.Conversion()
	osargs.Greeter()
}