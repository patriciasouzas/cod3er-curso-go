package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("linha.")

	fmt.Println("nova linha")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("o valor de x é" + xs)
	fmt.Println("o valor de x é", x)

	fmt.Printf("o valor de %.2f é", x)

	a := 6
	b := 1.968
	c := false
	d := "oi"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
