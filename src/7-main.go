package main

import "fmt"


func main() {

	switch module := 4 % 2; module {
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	// without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("It's greater than 100")
	case value < 0:
		fmt.Println("It's less than zero")
	default:
		fmt.Println("Value between 0 and 100")
	}
}
