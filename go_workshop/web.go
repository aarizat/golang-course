package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main(){
	response, _ := http.Get("http://google.com")
	e := webWriter{}
	io.Copy(e, response.Body)

}