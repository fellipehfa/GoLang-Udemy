package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { // não tem while no GOlang
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) // Sleep -> imprime, para por um segungo
		// caso queira maior espaço de tempo
		// time.Sleep(time.Second * 5)
	}
}
