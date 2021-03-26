package main

import "fmt"

func main() {
	// Declare contants
	const pi float64 = 3.14
	// other way - without declaring its type
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declare integer variables
	base := 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// zero values
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println(i, f, s, b)

	// square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square Area: ", squareArea)
}
