package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número inteiro: ")
	fmt.Scan(&num3)

	var menor int

	if num1 <= num2 && num1 <= num3 {
		menor = num1
	} else if num2 <= num1 && num2 <= num3 {
		menor = num2
	} else {
		menor = num3
	}

	fmt.Println("O menor número é:", menor)
}
