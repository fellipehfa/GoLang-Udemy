package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	fmt.Println("Executando!")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Não será transmitido")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
