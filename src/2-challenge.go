package main

import "fmt"

func rectArea (base, height float64) float64 {
	return base * height
}

func trapezeArea (base1, base2, height float64) float64 {
	return (base1 + base2) * height / 2
}

func circArea (radius float64) float64 {
	const pi = 3.1416
	return pi * radius * radius
}

func main() {
	rectangle := rectArea(10.0, 10.0)
	trapeze := trapezeArea(10.0, 10.0, 5.0)
	circle := circArea(1.0)
	fmt.Printf("Rectangle: %f \nTrapeze: %f \nCircle %f\n", rectangle, trapeze, circle)
}
