package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	// quando não coloca nenhum switch associado, ele procura por um case que bata com o solicitado
	// nesse caso ela é considerada true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
