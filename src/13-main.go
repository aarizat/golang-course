package main

import (
	"fmt"
	pkg "golang-course/src/mypackage"
)


func main() {
	var myCar pkg.CarPublic
	myCar.Brand = "Chevrolet"
	myCar.Year = 2030
	fmt.Println(myCar)

	pkg.PrintMsg("Hello World!")
	pkg.printMsg("Hello World!")

}