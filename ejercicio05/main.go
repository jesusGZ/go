package main

import "fmt"

func main() {
	/* i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	} */

	/* for i := 1; i <= 10; i++ {
		fmt.Println(i)
	} */

	/* num := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingreso el numero secreto:")
		fmt.Scanf("%d", &num)
		if num == 100 {
			break
		}
	} */

	var i = 0
	for i < 10 {
		fmt.Println("\n valor de i: %d", i)
		if i == 5 {
			fmt.Println("multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Println(i)
	}
}
