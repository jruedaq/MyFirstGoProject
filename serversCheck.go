package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	servers := []string {
		"https://www.oneago.com",
		"https://www.funmarfancol.org",
		"https://www.reintegra.org.co",
		"https://www.juandavid.dev",
		"https://noexist.juandavid.dev",
	}

	for _, server := range servers {
		checkServers(server)
	}

	totalTime := time.Since(start)

	fmt.Println()
	fmt.Println(totalTime, "hasta finalizar la ejecución")
}

func checkServers(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Println("El servidor", server, "no está disponible")
	} else {
		fmt.Println("El servidor", server, "funciona correctamente")
	}
}
