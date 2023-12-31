// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// conditions

	/*
		// var value bool
		var value string

		fmt.Print("Please enter a 'value': ")
		fmt.Scan(&value)

		if value == "false" {
			fmt.Println("false")
		} else if value == "true" {
			fmt.Println("true")
		} else {
			fmt.Println(value)
		}
	*/

	// switch
	// day := 1
	var day int

	fmt.Print("please enter a 'day': ")
	fmt.Scan(&day)

	switch day {
	case (1):
		fmt.Println("one")
	case (2):
		fmt.Println("two")
	case (3):
		fmt.Println("three")
	case (4):
		fmt.Println("four")
	case (5):
		fmt.Println("five")
	case (6):
		fmt.Println("six")
	case (7):
		fmt.Println("seven")
	default:
		fmt.Println(day)
	}

	// while loop functionality

	for day != 31 {
		fmt.Println("wrong") // infinite loop
	}

}

// In Go any executable code belongs to the main package.
// At the end of each line there is an implicit semicolon added so
// if you are working with any statement mind if you are creating new line
