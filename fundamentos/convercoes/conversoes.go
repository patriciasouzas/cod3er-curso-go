package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Test" + string(123))

	// int para string
	fmt.Println("Test" + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// bool para string
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("verdadeiro")
	}
}
