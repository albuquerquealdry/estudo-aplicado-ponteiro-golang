package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updateBookTitle(book *Book, newTitle string) {
	fmt.Println(&book.Title)
	book.Title = newTitle

}

func main() {
	myBook := Book{
		Title:  "Tommie",
		Author: "Junjiito",
		Pages:  300,
	}

	fmt.Println("Origin Title:", myBook.Title)
	fmt.Println(&myBook.Title)

	updateBookTitle(&myBook, "Horror Stories in Japan")

	fmt.Println("Updated Title:", myBook.Title)
}
