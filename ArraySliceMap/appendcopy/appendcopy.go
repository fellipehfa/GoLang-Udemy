package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(slice1, 4, 5, 6) -> array não recebe append

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // capacidade para 2 posições
	copy(slice2, slice1) //copy(x, y) -> x recebe, y é copiado (não pode ser um array)
	fmt.Println(slice2) // não altera a capacidade maxima prédefinida
}
