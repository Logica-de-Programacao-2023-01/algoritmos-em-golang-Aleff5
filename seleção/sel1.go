package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	var maior int

	if num1 > num2 {
		maior = num1
	} else {
		maior = num2
	}

	fmt.Println("O maior número é:", maior)
}
