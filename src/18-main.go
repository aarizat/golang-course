package main

import (
	"fmt"
)


func message(msg string, ch chan string) {
	ch <- msg
}

func main() {
	ch := make(chan string, 2)
	ch <- "Message1"
	ch <- "Message2"
	fmt.Println(len(ch), cap(ch))

	// close and range
	close(ch)
	//ch <- "Message3" // error

	for msg := range ch {
		fmt.Println(msg)
	}

	// select
	email1 :=  make(chan string)
	email2 :=  make(chan string)
	go message("Message 1", email1)
	go message("Message 2", email2)

	for i:= 0; i < 2; i++ {
		select {
		case m1 := <- email1:
			fmt.Println("Email got from email 1", m1)
		case m2 := <- email2:
			fmt.Println("Email got from email 2", m2)
		}
	}





}
