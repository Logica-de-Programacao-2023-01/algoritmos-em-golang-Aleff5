package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	if num1 >= 0 && num2 >= 0 {
		resultado := num1 * num2
		fmt.Printf("Resultado da multiplicação: %d\n", resultado)
	} else {
		resultado := num1 + num2
		fmt.Printf("Resultado da soma: %d\n", resultado)
	}
}
