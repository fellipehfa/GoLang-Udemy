package main

import (
	"fmt"

	"github.com/fellipehfa/go-modules/getlink"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - misturar (mensagens) em um cannal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		getlink.Titulo("https://www.cod3r.com.br", "http://www.google.com"),
		getlink.Titulo("http://www.amazon.com", "http://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
