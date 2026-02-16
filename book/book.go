package book

import "fmt"

type Book struct {
	Title string
	Autor string
	Pages int
}

func (b *Book) PrintInfo() {
	fmt.Println("Title: ", b.Title, "\n Author: ", b.Autor, "\n Pages: ", b.Pages)
}
