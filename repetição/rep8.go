package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&numero)

	fmt.Printf("Divisores do número %d:\n", numero)

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Println(i)
		}
	}
}
