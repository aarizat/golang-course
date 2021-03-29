package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	servers := []string {
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	start := time.Now()
	for _, server := range servers {
		verify_server(server)
	}
	t := time.Since(start)
	fmt.Printf("Time execution %s\n", t)
}

func verify_server(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Printf("Server %s not available\n", server)
	} else {
		fmt.Printf("Server %s available\n", server)
	}
}
