package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type webWriter struct {
	
}

func (webWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}

func main() {
	response, err := http.Get("http://google.com")

	if err != nil {
		log.Fatal(err)
	}

	e := webWriter{}

	io.Copy(e, response.Body)
}
