package main

import "fmt"

func main() {
	s := make([]int, 10) // criou um array inteiro com 10 elementos
	s[9] = 12            // insere um valor no indice 9
	fmt.Println(s)

	s = make([]int, 10, 20)        // cria um array da mesma forma, porem o array interno tem 20 elementos
	fmt.Println(s, len(s), cap(s)) // mostra o array s, o tamanho dele atualmente(quantos elementos já estão definidos) e sua capacidade

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // define os outros elementos ainda não preenchidos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // Adiciona mais um item ao array que já estava completo, por tanto ele expande a capacidade maxima do array, dobrando automaticamente sua capacidade
	fmt.Println(s, len(s), cap(s))
}
