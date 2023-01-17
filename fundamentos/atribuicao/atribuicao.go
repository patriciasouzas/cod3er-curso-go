package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	fmt.Println(i)

	// declaração de mais de uma variável + atribuição
	x, y := 1, 3
	fmt.Println(x, y)

	// substituição de valores
	x, y = y, x
	fmt.Println(x, y)
}
