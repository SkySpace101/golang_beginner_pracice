// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// struct definition
	type record struct {
		firstName string
		lastName  string
		salary    float64
		id        string
		role      string
	}

	// instantiation of the struct record type variable
	var r1 record

	r1.firstName = "shubham"
	r1.lastName = "Agarwal"
	r1.salary = 200000
	r1.id = "sad23ds"
	r1.role = "Senior Developer"

	fmt.Println(r1)
	fmt.Println(r1.salary)

	// map

	mymap := map[int]string{1: "game", 2: "sports", 3: "studies"}

	fmt.Println(mymap)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		// fmt.Println(mymap[i])
	}

	// getting input from the user

	var a, b int
	fmt.Println("Please enter a number")
	fmt.Scan(&a, &b)
	fmt.Println("thank you")

	fmt.Println("please enter a number")
	// fmt.Scanln(&b)
	fmt.Println(b)

	var c int
	fmt.Println("please enter a number")
	fmt.Scanf("%v", &c) // used when the input is required in a particular format specified in this Scanf function.
	fmt.Println("thanks")
	fmt.Println(c)

	fmt.Println(a)

	print_function()
	print_function()
	print_function()

	fmt.Println(add(31, 32))

}

func print_function() {
	fmt.Println("What is good?")
}

func add(a int, b int) (result int) {
	result = a + b
	return
}

// In Go any executable code belongs to the main package.
