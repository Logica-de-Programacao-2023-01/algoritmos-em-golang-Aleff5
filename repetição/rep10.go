package main

import "fmt"

func main() {
	var (
		numero, maior    int
		primeiraIteracao = true
	)

	fmt.Println("Digite os números inteiros (digite 0 para sair):")

	for {
		fmt.Print("Número: ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		if primeiraIteracao {
			maior = numero
			primeiraIteracao = false
		} else {
			if numero > maior {
				maior = numero
			}
		}
	}

	if primeiraIteracao {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		fmt.Printf("O maior número digitado é: %d\n", maior)
	}
}
