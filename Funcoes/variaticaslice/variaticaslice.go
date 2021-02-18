package main

import "fmt"

func imprimirAprovaods(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{ // não pode ser um array (numero de vetores não pode ser prédefinido)
		"Maria",
		"Pedro",
		"Rafael",
		"Guilherme",
	}
	imprimirAprovaods(aprovados...)
}
