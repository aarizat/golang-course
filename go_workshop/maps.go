package main

import "fmt"


func main() {
	m1 := make(map[string]int)

	m1["a"] = 28
	fmt.Println(m1)
	fmt.Println(m1["a"])
}