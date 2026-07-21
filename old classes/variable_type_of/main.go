package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		name      string
		countdown int     = 3
		action    bool    = true
		chance    float64 = 54.67
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("name: %s \ncountdown: %d \naction: %t \nchance: %.2f \n"+
		"Variables are of type (%T, %T, %v, %v) respectively\n",
		name, countdown, action, chance,
		name, countdown, reflect.TypeOf(action), reflect.TypeOf(chance))
	fmt.Printf("Participant name is %s ", name)
	fmt.Printf("at the count of %d you will take action. ", countdown)
	fmt.Printf("Your chance of success is %.2f%% .", chance)
}
