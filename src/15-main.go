package main

import "fmt"

// structs
type square struct {
	side float64
}

type rectangle struct {
	base float64
	heigth float64
}

// methods
func (s square) area() float64 {
	return s.side * s.side
}

func (r rectangle) area() float64 {
	return r.base * r.heigth
}


// interfaces
type figures2D interface {
	area() float64
}

func calculate(f figures2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	mySquare := square{10}
	myRectangle := rectangle{20, 20}
	// calling method
	fmt.Println(mySquare.area()) 
	fmt.Println(myRectangle.area())
	// Using interface
	calculate(mySquare)
	calculate(myRectangle)

	// List of interfaces
	myInterface := []interface{}{"Hi", 12, 4.55}
	fmt.Println(myInterface...)
}