package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123456] = "Maria"
	aprovados[147852] = "João"
	aprovados[369852] = "Pedro"
	aprovados[987456] = "Carlos"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[369852])
	delete(aprovados, 369852)
	fmt.Println(aprovados[369852])
}
