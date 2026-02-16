package main

import "githiub.com/DavidCT9/golangOOP/book"

func main() {
	// Golang doesn't implements OOP
	// Classes -> Structures
	// Methods -> Methods over structures
	// Encapsulation -> There no modifiers such things as private or public
	// If a function name starts with a capital letter, it means public, otherwise it is private

	// Inheritance -> Composition: Structures within structures
	// Polymorphism -> Interfaces

	var myBook = book.Book{
		Title:  "Los 5 lenguajes del amor",
		Author: "Gary Chapman",
		Pages:  200,
	}

	myBook.PrintInfo()

}
