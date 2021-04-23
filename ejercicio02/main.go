package main

import (
	"fmt"
	"strconv"
)

var Numero int
var texto string
var status bool

func main() {
	//var numero2, numero4, numero5 int
	numero3 := 4

	numero2, numero4, numero5, texto, status := 2, 5, 67, "variable texto", true

	var moneda float64 = 0

	numero2 = int(moneda)
	//texto = fmt.Sprintf("%d", moneda)
	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(texto)
	fmt.Println(status)

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
