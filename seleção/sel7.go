package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o salário do funcionário: ")
	fmt.Scanln(&salario)

	var novoSalario float64

	if salario < 1000.00 {
		novoSalario = salario * 1.1 // Aumento de 10%
	} else {
		novoSalario = salario * 1.05 // Aumento de 5%
	}

	fmt.Printf("Novo salário: R$ %.2f\n", novoSalario)
}
