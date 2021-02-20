package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	go fale("Maria", "Por que você não fala comigo?", 3)
	go fale("João", "Só posso falar depois de você!", 1)

	go fale("Maria", "ei", 500)
	go fale("João", "hey", 500)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabens", 5)
}
