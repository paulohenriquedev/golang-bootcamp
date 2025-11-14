package main

import (
	"fmt"
)

func main() {
	listaNotas := make([]float64, 0)
	var nota float64
	var notaTotal float64

	for {
		fmt.Println("Insira as notas!\nQuando terminar insira valor negativo")
		fmt.Println("Qual a nota do aluno?")
		fmt.Scan(&nota)
		if nota < 0 {
			break
		} else {
			listaNotas = append(listaNotas, nota)
			notaTotal += nota
		}
	}

	media := notaTotal / float64(len(listaNotas))

	fmt.Println("A média das notas é: ", media)

}
