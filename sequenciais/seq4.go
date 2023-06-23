package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	mediaPonderada := (num1*2 + num2*3 + num3*5) / (2 + 3 + 5)

	fmt.Printf("A média ponderada é: %.2f\n", mediaPonderada)
}
