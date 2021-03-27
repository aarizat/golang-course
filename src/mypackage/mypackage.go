package mypackage

import "fmt"

// CarPublic car with public access -- start with uppercase
type CarPublic struct {
	Brand  string
	Year int
}


// carPrivate with private access -- Start with lowercase
type carPrivate struct {
	brand string
	year int
}


// PrintMsg prints a message --> public access
func PrintMsg(msg string) {
	fmt.Println(msg)
}


// printMsg print a message --> private access
func printMsg(msg string) {
	fmt.Println(msg)
}
