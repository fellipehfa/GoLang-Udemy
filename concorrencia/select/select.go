package main

import (
	"fmt"
	"time"

	"github.com/fellipehfa/go-modules/getlink"
)

func fastest(url1, url2, url3 string) string {
	c1 := getlink.Titulo(url1)
	c2 := getlink.Titulo(url2)
	c3 := getlink.Titulo(url3)

	// estrutura de crontrole especifica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := fastest(
		"https://www.cod3r.com.br",
		"http://www.google.com",
		"http://www.amazon.com",
	)
	fmt.Println(campeao)
}
