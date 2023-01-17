package main

import "fmt"

func main() {
	i := 1
	for i <= 10 { // não tem while em go
		fmt.Printf("%d", i)
		i++
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}
}
