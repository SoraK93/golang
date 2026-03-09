package main

import (
	"fmt"
	"time"
)

func calculateSum(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()

	for i := 0; i < 10000; i++ {
		go calculateSum(i)
	}

	elapsed := time.Since(start)

	time.Sleep(2 * time.Second)
	fmt.Println("Function took", elapsed)
}
