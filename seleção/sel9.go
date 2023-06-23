package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite três números reais:")
	fmt.Print("Número 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Número 2: ")
	fmt.Scanln(&num2)
	fmt.Print("Número 3: ")
	fmt.Scanln(&num3)

	if num1 <= num2 && num1 <= num3 {
		if num2 <= num3 {
			fmt.Println("Números em ordem crescente:", num1, num2, num3)
		} else {
			fmt.Println("Números em ordem crescente:", num1, num3, num2)
		}
	} else if num2 <= num1 && num2 <= num3 {
		if num1 <= num3 {
			fmt.Println("Números em ordem crescente:", num2, num1, num3)
		} else {
			fmt.Println("Números em ordem crescente:", num2, num3, num1)
		}
	} else {
		if num1 <= num2 {
			fmt.Println("Números em ordem crescente:", num3, num1, num2)
		} else {
			fmt.Println("Números em ordem crescente:", num3, num2, num1)
		}
	}
}
