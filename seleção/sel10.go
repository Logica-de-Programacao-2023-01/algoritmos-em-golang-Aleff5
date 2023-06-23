package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade da pessoa: ")
	fmt.Scanln(&idade)

	if idade <= 9 {
		fmt.Println("Classificação: Mirim")
	} else if idade >= 10 && idade <= 13 {
		fmt.Println("Classificação: Infantil")
	} else if idade >= 14 && idade <= 17 {
		fmt.Println("Classificação: Juvenil")
	} else {
		fmt.Println("Classificação: Adulto")
	}
}
