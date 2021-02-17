package main

import "fmt"

func main() {
	// são estruturas fixas, ou seja, contem o mesmo tipo de dado e é estático (numero fixo de arrays), caso defina um numero especifico de arrays, ele sempre terá esse mesmo numero
	var notas [3]float64 // [3] quantidade de arrays \\ float64 tipo estático
	fmt.Println(notas)
	// Por padrão, se não for definido um valor para os arrays, automaticamento ele recebe o valor 0

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ { // len -> tamanho do array ou characters
		total += notas[i]
	}

	media := total / float64(len(notas)) // transforma o valor de inteiro para float64
	fmt.Printf("Média %.2f\n", media)
}
