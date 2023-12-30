// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// mutiple variable declaration and assignment

	// on a single line
	// var var1, var2, var3 float64
	var var1, var2, var3 float64 = 34.5, 56.4, 43.3434

	// different types of variables in single line
	var numVar, fltVar, strVar, boolVar = 2, 45.7, "who is this", true

	// different types using walrus on same line
	numVar1, strVar1, boolVar1 := 4, "ashfahsfb", true

	// much better declaration and assignment of variables
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	fmt.Println(numVar)
	fmt.Println(fltVar)
	fmt.Println(strVar)
	fmt.Println(boolVar)

	fmt.Println(numVar1)
	fmt.Println(strVar1)
	fmt.Println(boolVar1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Constant Declaration and Assignment - value must be assigned when a constant is declared.

	const CONSTANTNUM int = 30 // typed constant

	fmt.Println(CONSTANTNUM)

	const X_VALUE = 34.5 // untyped constant

	fmt.Println(X_VALUE)

	// X_VALUE = 45  // ERROR: Cannot assign to a Constant

	const (
		Y_VALUE         = 23
		Z_VALUE float32 = 34.66
		K_VALUE         = "GOOP"
		P_VALUE         = false
	)

	fmt.Println(Y_VALUE)
	fmt.Println(Z_VALUE)
	fmt.Println(K_VALUE)
	fmt.Println(P_VALUE)

}

// In go any executable code belongs to the main package.
