// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	fmt.Print("hello")
	fmt.Print("World")

	// using formatting verbs in printf
	fmt.Printf("This is a fund %d and %5.3f and  %s %x  %T,%T,%T %v %v %v %t %e %g", 45, 343.435, "shubham", 34, "sdfsd", 34, 34.56, "sdfsd", 34, 34.56, false, 56.45, 67.1345454)

}

// In Go any executable code belongs to the main package.
