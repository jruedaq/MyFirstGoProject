package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese una suma con 2 operadores")
	scanner.Scan()
	operation := scanner.Text()

	operator := "/"

	values := strings.Split(operation, operator)

	factor1, e1 := strconv.Atoi(values[0])
	if e1 != nil {
		log.Fatal("El primer parámetro falló\n", e1)
	}
	factor2, e2 := strconv.Atoi(values[1])
	if e2 != nil {
		log.Fatal("El segundo parámetro falló\n", e2)
	}

	var result float64
	switch operator {
	case "+":
		result = float64(factor1 + factor2)
		break

	case "-":
		result = float64(factor1 - factor2)
		break

	case "*":
		result = float64(factor1 * factor2)
		break

	case "/":
		result = float64(factor1 / factor2)
		break

	default:
		log.Fatal("No se detectó un operador válido")
	}

	fmt.Println("El resultado es", result)
}
