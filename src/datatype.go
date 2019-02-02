package main

import "fmt"

/*
basic data types:
Integers:
uint, uint8, uint16, uint 32, uint 64
int, int8, int16, int32, int64

Float:
float32, float64

Complex:
complex64, complex128

Boolean:
bool(true, false)

String:
string

nil
*/

func main() {
	// array
	// var days []string, or
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// slice
	weekdays := days[0:5]
	fmt.Println(weekdays)

	// map
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}
	fmt.Println(youtubeSubscribers)

	// struct
	// our Person struct
	type Person struct {
		name string
		age  int
	}

	// declaring a new `Person`
	var myPerson Person
	fmt.Println(myPerson)

	// declaring a new `elliot`
	elliot := Person{name: "Elliot", age: 24}
	// trying to roll back time to before I was injury prone
	elliot.age = 18

	// our Boss struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Boss'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
	// declaring a boss with employees
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
