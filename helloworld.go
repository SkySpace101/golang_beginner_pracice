// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	/*
			// Array declaration and Assignment.
			var arr = [4]int{1, 2, 3, 4}

			fmt.Println(arr)

			// Declaring and Assigning using walrus operator
			arr1 := [5]string{"a", "b", "c"} // rest unassigned elements carry default value for that datatype.

			fmt.Println(arr1)

			// arrays with inferred lengths
			arr2 := [...]int{1, 2, 3, 4, 5, 5, 6}

			fmt.Println(arr2)

			// accessing the elements of an array
			fmt.Println(arr2[5])
			fmt.Println(arr2[3])

			// changing the value of an element of an array
			arr1[0] = "k"
			arr2[5] = 41

			fmt.Println(arr2)

			// printing an out of bounds element
			// fmt.Println(arr1[10]) -- Error: Out of bounds



		arr1 := [5]int{}              //not initialized
		arr2 := [5]int{1, 2}          //partially initialized
		arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(arr3)

	*/

	// Initialize only Specific elements
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)

	// getting the length of an array

	fmt.Println(len(arr1))
}

// In Go any executable code belongs to the main package.
