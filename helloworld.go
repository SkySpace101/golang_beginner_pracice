// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// variable declaration method - 1 - using var
	var integer int // default 0
	var flt float32 // default 0
	var str string  // default ""
	var bo bool     // default false

	fmt.Println(integer)
	fmt.Println(flt)
	fmt.Println(str)
	fmt.Println(bo)
	fmt.Println()

	integer = 23  // assigning value to the int variable
	flt = 34.6    // assigning value to the float variable
	str = "value" // assigning value to the string variable
	bo = true     // assigning value to the bool variable

	fmt.Println(integer)
	fmt.Println(flt)
	fmt.Println(str)
	fmt.Println(bo)
	fmt.Println()

	// Variable declaration and Inferring type from assignment. (using Walrus (:=) operator)

	intNum := 34
	fltNum := 56.546
	str1 := "tree"
	bol := true

	fmt.Println(intNum)
	fmt.Println(fltNum)
	fmt.Println(str1)
	fmt.Println(bol)

}

// In go any executable code belongs to the main package.
