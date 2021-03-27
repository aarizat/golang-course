package main

import "fmt"


func main() {
	// for conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	// for while
	counter := 0
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	// infinite loop
	j := 0
	for {
		println(j)
		j++
	}

}