package main

import "fmt"

func main() {
	var (
		numero     int
		soma       int
		quantidade int
	)

	fmt.Println("Digite os números inteiros (digite 0 para sair):")

	for {
		fmt.Print("Número: ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		quantidade++
	}

	if quantidade == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		media := float64(soma) / float64(quantidade)
		fmt.Printf("A média aritmética é: %.2f\n", media)
	}
}
