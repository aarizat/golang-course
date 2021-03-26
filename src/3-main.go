package main

import "fmt"

func main() {
	// Println
	firstMessage := "Hello"
	secondMessage := "World!"
	fmt.Println(firstMessage, secondMessage)
	fmt.Println(secondMessage)
	// Printf
	name := "Andres"
	age := 28
	fmt.Printf("%s is %d year old\n", name, age)
	fmt.Printf("%v is %v year old\n", name, age) // without specifyng type
	// Sprinf
	message := fmt.Sprintf("%s is %d year old", name, age)
	fmt.Println(message)
	// printing types of variables
	fmt.Printf("name %T\n", name)
	fmt.Printf("age %T\n", age)
}