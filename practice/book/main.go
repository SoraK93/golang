package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
	(*book).Pages = pages
}

func main() {
	books := []*Book{
		{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Pages:  180,
		},
		{
			Title:  "To Kill a Mockingbird",
			Author: "Harper Lee",
			Pages:  281,
		},
		{
			Title:  "Pride and Prejudice",
			Author: "Jane Austen",
			Pages:  279,
		},
	}

	pageCount := []int{210, 250, 295}
	
	for index, book := range books {
		updatePages(book, pageCount[index])
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
