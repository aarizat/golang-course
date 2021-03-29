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
	channel := make(chan string)

	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range servers {
			go verify_server(server, channel)
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-channel)
		i++
	}

	t := time.Since(start)
	fmt.Printf("Time execution %s\n", t)
}

func verify_server(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " It's not available"
	} else {
		channel <- server + " It's available"
	}
}
