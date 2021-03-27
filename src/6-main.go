package main

import "fmt"


func main(){
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("It's 1")
	} else {
		fmt.Println("It's not 1")
	}

	// with and
	if value1 == 1 && value2 == 2 {
		fmt.Println("They are 1 and 2")
	}

	// with or
	if value1 == 0 || value2 == 2 {
		fmt.Println("There is one true")
	}

}