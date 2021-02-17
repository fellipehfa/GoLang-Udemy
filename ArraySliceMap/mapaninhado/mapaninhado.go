package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"B": {
			"Baiana": 15456.78,
			"Bruno":  154654.87,
		},
		"R": {
			"Rael": 8456.78,
		},
		"F": {
			"Fellipe": 1200.0,
		},
	}
	delete(funcPorLetra, "F")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
