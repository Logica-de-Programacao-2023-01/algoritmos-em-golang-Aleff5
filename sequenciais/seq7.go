package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	sucessor := numero + 1
	antecessor := numero - 1

	fmt.Println("O sucessor do número é:", sucessor)
	fmt.Println("O antecessor do número é:", antecessor)
}
