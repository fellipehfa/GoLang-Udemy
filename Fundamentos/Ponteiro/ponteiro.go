package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritiméticas de ponteiros
	var p *int = nil

	p = &i // & -> pega o endereço da variavel i
	*p++   //*p -> pega o valor associado ao ponteiro
	// desreferenciando para pegar o valor
	i++

	fmt.Println(p, *p, i, &i)
}
