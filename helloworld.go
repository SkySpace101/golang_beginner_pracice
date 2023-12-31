// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// Four basic types of datatype in Go.
	// 1. Integer
	// 2. Float
	// 3. String
	// 4. Boolean

	var a int64
	a = 23

	fmt.Println(a)

	// var b int8 -- produce an error if try to assign 200
	var b uint8
	b = 200

	fmt.Println(b)

	var f1 float32
	var f2 float64

}

// In Go any executable code belongs to the main package.
