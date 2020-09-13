package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	channel := make(chan string)

	servers := []string {
		"https://www.oneago.com",
		"https://www.funmarfancol.org",
		"https://www.reintegra.org.co",
		"https://www.juandavid.dev",
		"https://noexist.juandavid.dev",
	}

	i := 0
	for {

		if i >= 10 {
			break
		}

		for _, server := range servers {
			go checkServers(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)

		i++
	}
}

func checkServers(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		channel <- "El servidor " + server + " no está disponible"
	} else {
		channel <- "El servidor " + server + " funciona correctamente"
	}
}