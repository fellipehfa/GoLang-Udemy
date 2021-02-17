package main

import "fmt"

// Não possio operadores ternarios
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	// return nora >= 6 ? "Aprovado" -> Metodo Javascript
}

func main() {
	fmt.Println(obterResultado(5))
}
