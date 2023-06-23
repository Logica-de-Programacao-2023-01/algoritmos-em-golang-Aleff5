package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	fmt.Printf("Tabuada de multiplicação para o número %d:\n", numero)

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
