package main

import "fmt"

func main() {
	var peso, altura float64

	fmt.Print("Digite o peso (em kg): ")
	fmt.Scan(&peso)

	fmt.Print("Digite a altura (em metros): ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("O IMC é: %.2f\n", imc)
}
