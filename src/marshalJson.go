package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	// an instance of our Book struct
	book := Book{Title: "Learning Concurreny in Python", Author: "Elliot Forbes"}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	type Author struct {
		Sales     int  `json:"book_sales"`
		Age       int  `json:"age"`
		Developer bool `json:"is_developer"`
	}

	type Book1 struct {
		Title  string `json:"title"`
		Author Author `json:"author"`
	}

	author := Author{Sales: 3, Age: 25, Developer: true}
	book1 := Book1{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err = json.Marshal(book1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	byteArray, err = json.MarshalIndent(book1, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
