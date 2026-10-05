package main

import ("fmt"; "path")

/*
Path package provides utility functions for working with url path strings
path.Split()
func Split(path string) (dir, file string)
path = assets/agreements.pdf
dir = assets/
file = agreements.pdf
*/

func PathSeparator() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir :", dir)
	fmt.Println("file :", file)

	
	// if dir not required, we can skip it using "_"
	var file1 string
	_, file1 = path.Split("css/main1.css")
	fmt.Println("file1 :", file1)

	// trying short declaration to directly assign return values to a variable
	_, file2 := path.Split("css/main2.css")
	fmt.Println("file2 :", file2)
}