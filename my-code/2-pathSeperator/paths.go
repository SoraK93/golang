package pathseperator

// path package: provide utility functions for working with url paths strings
// func Split (path string) (dir, file string)

import ("path")

func Path() (string) {
	var file string
	_, file = path.Split("css/styles.css")
	
	// short declaration statement
	// _, file := path.Split("css/styles.css")

	return file
}