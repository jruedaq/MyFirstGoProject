package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware{
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			flag := true
			fmt.Fprintf(writer, "Checking Auth!")

			if flag {
				handlerFunc(writer, request)
			} else {
				return
			}
		}
	}
}

func Login() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(start))
			}()

			handlerFunc(writer, request)
		}
	}
}
