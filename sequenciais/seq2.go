package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	dobro := numero * 2
	triplo := numero * 3
	quadruplo := numero * 4

	fmt.Println("O dobro do número é:", dobro)
	fmt.Println("O triplo do número é:", triplo)
	fmt.Println("O quadruplo do número é:", quadruplo)

}
