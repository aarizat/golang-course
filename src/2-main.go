package main

import "fmt"

func main() {
	// arithmetic operators
	x := 10
	y := 50

	// sum
	result := x + y
	fmt.Println("The sum is :", result)

	// substract
	result = x - y
	fmt.Println("x minus y is :", result)

	// multiply
	result = x * y
	fmt.Println("x by y is:", result)

	// divide
	result = x / y
	fmt.Println("x divide y is:", result)

	// module
	result = x % y
	fmt.Println("x module y is:", result)

	// increment
	x++
	fmt.Println("x plus 1:", x)

	x--
	fmt.Println("x minus 1:", x)
}