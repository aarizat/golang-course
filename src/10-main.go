package main

import (
	"fmt"
	"strings"
)

func isPalindrome (s string) bool {
	steps, k, i := len(s) / 2, len(s) -1, 0
	for ; i < steps; i++ {
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[k])) {
			break
		}
		k--
	}
	return i == steps
}


func main() {
	slice := []string{"hi", "What", "do", "you", "do?"}
	for ix, val := range slice {
		fmt.Println(ix, val)
	}
	result := isPalindrome("Ana")
	fmt.Println(result)
}
