package main

import "fmt"

func main() {

	var idadeAnos int

	fmt.Print("Digite a idade em anos: ")
	fmt.Scan(&idadeAnos)

	idadeDias := idadeAnos * 365

	fmt.Println("A idade em dias é:", idadeDias)
}
