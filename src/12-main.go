package main

import "fmt"


type car struct {
	brand string
	year int
}

func main() {
	// instantiation
	myCar := car{brand: "Ford", year: 2030}
	fmt.Println(myCar)

	// other way
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar) // its year has zero value
}
