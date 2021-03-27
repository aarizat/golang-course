package main

import "fmt"


func main() {
	// maps
	m := make(map[string]int)
	fmt.Println(m)
	m["Andres"] = 28
	m["Carlos"] = 30
	fmt.Println(m)
	// iterate map
	for i, v := range m {
		fmt.Println(i, v)
	}
	// find values
	value := m["Andres"]
	fmt.Println(value)
	// When the value is not found
	fmt.Println(m["Felipe"]) // print its zero value according to type
	val, exist := m["Miguel"]
	fmt.Println(val, exist) // not exist --> 0 false
	val, exist = m["Andres"]
	fmt.Println(val, exist) // exist --> 2 true
}
