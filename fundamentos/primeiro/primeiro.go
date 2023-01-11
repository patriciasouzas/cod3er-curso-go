// programas executáveis iniciam pelo pacote main

package main

/*
os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/

import "fmt"

// a porta de entrada de um programa Go é a função main

func main() {
	fmt.Println("olá, mundo!")
}
