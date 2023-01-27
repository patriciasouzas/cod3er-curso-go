package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Pq você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)

	//fmt.Println("Fim!")

	go fale("Maria", "Oi", 10)
	fale("João", "Fala", 5)
}
