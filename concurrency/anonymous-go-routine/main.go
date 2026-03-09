package main

import (
	"fmt"
	"time"
)

func main() {
	func () {
		fmt.Println("In anonymous")
	}()
	time.Sleep(1 * time.Second)
}
