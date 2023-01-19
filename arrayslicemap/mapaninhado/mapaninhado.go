package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 14578,
			"Gustavo": 457854,
		},
		"J": {
			"João": 14578,
		},
		"P": {
			"Pedro": 457845,
			"Pablo": 4578451,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
