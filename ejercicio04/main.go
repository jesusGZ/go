package main

import "fmt"

var numero1 int
var numero2 int

func main() {
	fmt.Println("---- Suma ----")
	fmt.Println("Ingrese numero 1:")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Ingrese numero 2:")
	fmt.Scanf("%d", &numero2)

	fmt.Println("Resultado: ", numero1+numero2)

}
