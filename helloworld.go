// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// Slices - like arrays store similar types values contiguously.
	// ond difference that arrays are fixed size and slices can have variable size
	// - that means it can grow and shrink.

	// declaring a slice using []datatype{values}
	slice_name := []int{1, 2, 3, 4}

	fmt.Println(slice_name)

	// common way of decalring a slice
	// myslice := []int{} // empty slice
	// myslice := []int{1, 2, 3}

	// getting length
	// fmt.Println(len(myslice))

	// getting capacity
	// fmt.Println(cap(myslice))

	// fmt.Println(myslice)

	// Creating a Slice from an Array
	// var myarray = [length]datatype{values}
	// myslice := myarray[start:end] // a slice made from an array

	// arr1 := [6]int{10, 11, 12, 13, 14, 15}
	// myslice := arr1[2:4]

	// fmt.Printf("myslice = %v\n", myslice)
	// fmt.Printf("length = %d\n", len(myslice))
	// fmt.Printf("capacity = %d\n", cap(myslice)) // the capacity of slice is start index to end of 'array'.

	// Create a slice with make() function
	// slice_name := make([]type, Length, capacity)  - Syntax - if capacity not defined will be equal to length.

	// myslice1 := make([]int, 5, 10)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	// myslice2 := make([]int, 5)
	// fmt.Printf("myslice2 = %v\n", myslice2)
	// fmt.Printf("length = %d\n", len(myslice2))
	// fmt.Printf("capacity = %d\n", cap(myslice2))

	// Go Access, Change, Append and Copy Slices

	// Access
	// prices := []int{10, 20, 30}

	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// changing element of the slice
	// prices := []int{10, 20, 30}
	// prices[2] = 50
	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// Append Elements to the slice
	// slice_name = append(slice_name, element1, element2, ...) - Syntax

	// myslice1 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// myslice1 = append(myslice1, 20, 21)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// Append one slice to another
	// slice3 = append(slice1, slice2...) - Syntax
	// myslice1 := []int{1, 2, 3}
	// myslice2 := []int{4, 5, 6}
	// myslice3 := append(myslice1, myslice2...)

	// fmt.Printf("myslice3=%v\n", myslice3)
	// fmt.Printf("length=%d\n", len(myslice3))
	// fmt.Printf("capacity=%d\n", cap(myslice3))

	// Change The Length of a Slice

	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

}

// In Go any executable code belongs to the main package.
