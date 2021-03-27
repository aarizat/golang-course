package main

import "fmt"


func say(text string, c chan<- string) {
	c <- text
}

func main() {
	ch := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", ch)

	fmt.Println(<-ch)
}