// Package Declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions
func main() {
	fmt.Println("Hello World")

	// myslice := []int{31, 34, 45, 243, 324, 2343}
	// myslice1 := []int{234, 45, 2, 67, 34, 34}

	// myslice2 := append(myslice, myslice1...)

	mymap := map[int]string{}

	mymap[1] = "one"
	mymap[2] = "two"
	mymap[3] = "three"
	mymap[4] = "four"

	fmt.Println(mymap)

	for key, val := range mymap {
		fmt.Printf("the key is %d and the value is %s.\n", key, val)
	}

	// fmt.Println(myslice)
	// fmt.Println(myslice2)

	// for i := 0; i < len(myslice); i++ {
	// 	fmt.Println(myslice[i])
	// }
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))

}

// In Go any executable code belongs to the main package.
