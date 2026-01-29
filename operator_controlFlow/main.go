package main

import "fmt"

func main() {
	fmt.Println("\nComparison Operator")
	comparison()

	fmt.Println("\nArithematic Operator")
	arithematic()

	fmt.Println("\nLogical Operator")
	logical()

	fmt.Println("\nBitwise Operator")
	bitwise()

	fmt.Println("\nIf-else statement")
	ifElse()

	fmt.Println("\nSwitch statement")
	switchStatement()

	fmt.Println("\nLoop statement")
	loop()
}

func comparison() {
	var a int = 12
	var b int = 12
	fmt.Println(a <= b)

	a = 20
	fmt.Println(a <= b)

	b = 100
	fmt.Println(a <= b)

	c := 0
	fmt.Println(a <= c)
}

func arithematic() {
	var (
		a       float64 = 5.9
		b, c, d int     = 2, 10, 4
	)

	fmt.Println((int(a) + b) * c / d)
}

func logical() {
	var a bool = true
	result := 10 > 50
	fmt.Println(!(a || result))
}

func bitwise() {
	var x, y int = 100, 90
	fmt.Println(x & y)
	fmt.Println(x | y)
	fmt.Println((!(((x + y) >> 2) == 47)))
}

func ifElse() {
	var a, b string = "foo", "bar"
	if a+b == "foo" {
		fmt.Println("foo")
	} else if a+b == "foobar" {
		fmt.Println("bar")
	} else if a+b == "bar" {
		fmt.Println("foobar")
	} else {
		fmt.Println("None matched")
	}
	fmt.Println("thank you!")

}

func switchStatement() {
	var i, j = 10, 50
	switch {
	case i+j == 60:
		fmt.Println("equal to 60")
	case i+j <= 60:
		fmt.Println("less than or equal to 60")
		fallthrough
	default:
		fmt.Println("greater than 60")
	}

	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}
}

func loop() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i * i)
		if i == 3 {
			break
		}
	}
}
