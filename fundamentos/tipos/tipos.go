package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos) - uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "olá, meu nome é jhony"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é", len(s1))

	// string com múltiplas linhas
	s2 := `olá,
	meu
	nome
	é sandra`
	fmt.Println(s2)
	fmt.Println(len(s2))
}
