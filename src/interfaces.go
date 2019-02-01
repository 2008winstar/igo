package main

import "fmt"

//Defining Interfaces
/*
 If we define a type based off this interface then we are forced to implement all of the functions or methods defined
 within that interface type.
*/
type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return 23
}

func (e Engineer) Random() (string, error) {
	return "ok", nil
}

func main() {
	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)

	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
	fmt.Println(programmers)
}
