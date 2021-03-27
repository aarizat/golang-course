package main

import "fmt"


func main() {
	// defer
	defer fmt.Println("Hello")
	fmt.Println("World")

	// break and continue
	for i := 0; i < 10; i++ {

		fmt.Println(i)
		if i == 2 {
			fmt.Println("It's two")
			continue
		}

		if i == 8 {
			fmt.Println("Loop finished!")
			break
		}
	}
}