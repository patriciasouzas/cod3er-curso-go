package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      12456,
		"Gabriel Santos": 45785,
		"Maria José":     457852,
	}

	funcESalarios["Rafael"] = 14578
	delete(funcESalarios, "inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
