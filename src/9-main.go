package main

import "fmt"

func main() {
	// array
	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, len(arr), cap(arr))

	// slice
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	// append
	slice = append(slice, 7)
	fmt.Println(slice)

	// add another slice at the end
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
	slice = append(slice, 11, 12, 13)
	fmt.Println(slice)
}

