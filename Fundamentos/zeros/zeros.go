package main

import "fmt"

func main() {
	var a int     // resposta -> 0
	var b float64 // resposta -> 0
	var c bool    // resposta -> falso
	var d string  // resposta -> vazia ou ""
	var e *int    // resposta -> <nil> ou null

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
