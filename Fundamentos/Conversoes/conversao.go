package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // é necessário que as variaveis sejam explicitas para executar as formulas

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)
	// neste caso, como "notaFinal" foi declarada como inteiro, 6,9 será arredondado sempre para o menor valor

	// par concatenar, cuidado....
	fmt.Println("Teste" + string(123))
	fmt.Println("Teste" + string(97))
	// dessa forma você só faz referencia a uma tabela de strings

	// int para string
	fmt.Println("Teste" + strconv.Itoa(123)) // Itoa -> int to a

	// string par int
	num, _ := strconv.Atoi("123") // Atoi -> string to int
	// num, _ -> num é o valor que nos interessa e o segundo valor está sendo ignorado, isso é representado por "_"
	fmt.Println(num - 122)

	// bool to string
	b, _ := strconv.ParseBool("true") // valor verdadeiro tambem pode ser representado pelo numero "1", assim como valor falso é representado pelo numero "0"
	if b {
		fmt.Println("Verdadeiro")
	}
}
