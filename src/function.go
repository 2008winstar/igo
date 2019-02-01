package main

import "fmt"

/*
func name(parameter_list parameter_type) (result_list result_type)  {
	body
}
*/

func add(a, b int) int {
	return a + b
}

func myfunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func addOne() func() int {
	var x int
	// Anonymous Functions
	return func() int {
		x++
		return x + 1
	}
}

func main() {
	fmt.Println(add(3, 5))

	fullName, err := myfunction("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)

	myFunc := addOne()
	fmt.Println(myFunc()) // 2
}
