package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice -> tamanho variavel
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	// slice não é um array! Slice define um pedaço de um array
	s2 := a2[1:3] // pega o que estiver entre a posição 1 a 3, sem considerar a posição 3
	fmt.Println(a2, s2)
	// novo slice, mas aponta para o mesmo array
	s3 := a2[:2] // pega ate a posição 2, sem considerar a posição 2
	fmt.Println(a2, s3)
	// slice de um slice
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
