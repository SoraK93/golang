package basics

func ShortVariables() (int, int, string) {
	// Use this when you know the values of the variable
	// Use this to keep the code concise and easy to read
	width, height := 100, 50

	// Here we are reassigning and declaring a new variable at the same time. 
	width, height, color := 50, 75, "red"

	return width, height, color
}
