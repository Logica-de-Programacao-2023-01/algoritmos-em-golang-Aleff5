package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Printf("O número %d é múltiplo de 3 e de 5.\n", numero)
	} else {
		fmt.Printf("O número %d não é múltiplo de 3 e de 5 ao mesmo tempo.\n", numero)
	}
}
