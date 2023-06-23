package main

import "fmt"

func main() {
	var preco float64

	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&preco)

	desconto := preco * 0.1
	valorComDesconto := preco - desconto

	fmt.Printf("O valor com desconto de 10%% é: R$ %.2f\n", valorComDesconto)
}
