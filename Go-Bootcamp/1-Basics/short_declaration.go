package main

/*
When to use variable declaration vs short declaration

Short Declaration has many useful properties:
Example it can be used inside if and switch statements to create variables that belong to those statements only.
*/

/*
Use the normal/variable declaration when you need a package scoped variable
version := 0 // Dont work
*/
// var version string

func Declaration() {
	/* 
	Use Variable declaration
	If you don't know the initial value
	score := 0 // DON'T!!!!
	*/
	var score int // defaults to score = 0
	_ = score

	/* 
	Use variable declaration
	When you want to group variables together for greater readability
	*/
	var (
		// related:
		video string
		// closely related:
		duration int
		current int
	)
	_ = video
	_ = duration
	_ = current

	/*
	Use Short declaration
	If you know the initial value
	var width, height = 100, 50 // DON'T!!
	*/
	width, height := 100, 50
	_ = width
	_ = height

	/*
	Use Short declaration
	To keep the code concise and easy to read
	*/
	
	/*
	Use Short declaration
	for redeclaration (its both a blessing and a curse)
	DON'T DO THIS!!
	width = 50 // assigns 50 to width
	color := "red" // new variable: color

	Example: "https://go.dev/play/p/Fmn-s56YzVN"
	*/
	width, color := 50, "red"
	_ = color

}