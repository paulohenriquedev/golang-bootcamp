package main

import "fmt"

func main() {
	var numero int
	var ePrimo bool = true

	fmt.Println("Entre com 1 números e te digo se é primo: ")
	fmt.Scan(&numero)

	if numero < 2 {
		fmt.Println("Não é primo!")
		return // finaliza o programa
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			ePrimo = false
			break
		}
	}

	if ePrimo == false {
		fmt.Println("Não é primo!")
	} else {
		fmt.Println("É primo!")
	}
}
