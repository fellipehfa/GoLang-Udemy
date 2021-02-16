package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255 //"byte" refere-se a uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O valor maximo do int é", i2)
	fmt.Println("O rune é", reflect.TypeOf(i2))

	// numeros reais (float32 e float64)
	var x float32 = 49.99 // ou pode declarar dessa formar "var x = float32(49.99)"
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // Define por padrão float64 por causa do sistema operacional

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) //"!" contrario

	// String
	s1 := "Hello World"    // não pode ser representado por aspas simples ''
	fmt.Println(s1 + "!!") // concatenação
	fmt.Println("O tamanho da string é", len(s1))

	// String com multiplas linhas
	s2 := `Hello, 
	name 
	is 
	Fellipe`
	fmt.Println(s2)
	fmt.Println("O tamanho da istring é", len(s2))

	// char???
	char := 'a'
	// var x char = 'b' -> não existe variavel tipo "char" na linguagem
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
	// dessa forma ela é reconhecida com int32
}
