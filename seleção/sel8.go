package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número entre 1 e 7: ")
	fmt.Scanln(&numero)

	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	fmt.Println("Dia da semana:", diaDaSemana)
}
