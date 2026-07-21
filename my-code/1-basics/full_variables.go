package basics

// This variable is scoped to this package.
var package_scope string

func FullVariables() int {
	// Use it if you do not know the initial value
	var new_variable int = 1234

	// Use this declaration only for closely related variables.
	// var (
	// 	// related
	// 	video string

	// 	// closely related
	// 	duration int
	// 	current int
	// )

	return new_variable
}