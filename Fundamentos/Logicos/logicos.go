package main

import "fmt"

func compras(work1, work2 bool) (bool, bool, bool) {
	comprarTv50 := work1 && work2
	comprarTv32 := work1 != work2 // ou exclusivo
	comprarSorvete := work1 || work2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
