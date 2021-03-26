package main

import "fmt"

func main() {
	// calculate rectangle, trapeze and circle
	// Rectangle area
	baseRect := 10
	heightRect := 10
	areaRect := baseRect * heightRect
	fmt.Println(areaRect)
	// Trapeze area
	base1 := 10
	base2 := 20
	height := 10
	areaTrapeze := (base1 + base2) * height / 2
	fmt.Println(areaTrapeze)
	// circle area
	const pi = 3.1416
	radius := 1.0
	area := pi * radius * radius
	fmt.Println(area)
}
