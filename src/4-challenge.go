package main

import "fmt"


func isPar (a int) bool {
	return a % 2 == 0
}

func verifyUser (username, pswd string) bool {
	return username == "aarizat" && pswd == "12345"
}

func main() {
	p := isPar(10)
	u := verifyUser("aarizat", "1234")
	fmt.Printf("p is par: %v\n", p)
	fmt.Printf("the user is autorized: %v\n", u)
}
