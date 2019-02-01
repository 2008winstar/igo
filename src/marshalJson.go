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
}
