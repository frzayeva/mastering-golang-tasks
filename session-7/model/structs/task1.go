package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func main() {

	book1 := Book{

		title:  "The Go Programming Language",
		author: "Alan A. A. Donovan",
		pages:  380,
	}
	fmt.Println(book1) // why ln is not working?(
	//fmt.Printf("%T", book1)

}