package main

import (
	. "10.struct/test"
	"fmt"
)

type myInt int

func changeBook(book Book) {
	book.Author = "111"
}

func changeBook2(book *Book) {
	book.Author = "111"
}

func main() {

	var a myInt = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var book = Book{Author: "Alen", Title: "How to study Java"}
	book.Title = "1"
	fmt.Printf("%v\n", book)
	fmt.Printf("%T\n", book)

	changeBook(book)
	fmt.Printf("%v\n", book)

	changeBook2(&book)
	fmt.Printf("%v\n", book)

	book.SetTitle("sss")
	fmt.Printf("%v\n", book)

}
