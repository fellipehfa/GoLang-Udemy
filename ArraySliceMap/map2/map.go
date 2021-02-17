package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Baiana":  11325.45,
		"Rael":    15456.78,
		"Fellipe": 1200.0, // ultimo elemento tambem precisa de virgula
	}

	funcsESalarios["Marcello"] = 1350.0 // adiciona mais um vetor
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
