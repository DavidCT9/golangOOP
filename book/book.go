package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Println("Title: ", b.Title, "\n Author: ", b.Author, "\n Pages: ", b.Pages)
}
