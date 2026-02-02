package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	var testList = strings.Split(text, " ")

	textMap := map[string]int{}

	for _, value := range testList {
		if (textMap[value] > 0) {
			textMap[value] += 1
		} else {
			textMap[value] = 1
		}
	}
	
	return textMap
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}