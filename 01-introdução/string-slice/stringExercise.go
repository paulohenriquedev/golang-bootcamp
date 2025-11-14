package main

import (
	"fmt"
	"slices"
)

func main() {
	var array [5]string
	array[0] = "paulo"
	array[1] = "henrique"
	array[2] = "da"
	array[3] = "silva"

	fmt.Println(array)
	fmt.Println(array[0:1], array[2:4])
	array[4] = "santos"
	fmt.Println(array, "  O tamanho do array é: ", len(array))

	//teste de slice

	slice := make([]string, 5)

	slice[0] = "Paulo"
	slice[1] = "Henrique"
	// os indices 2,3,4 estão vazios pois foi definido 5 posições
	//add mais 3 itens
	slice = append(slice, "da", "silva", "santos")
	//print do slice agora com as novas posições
	fmt.Println(slice, "tamanho do slice: ", len(slice))
	// removendo o espaço em branco 2,3 e 4
	indiceRemover := 2
	slice = slices.Delete(slice, indiceRemover, indiceRemover+3)
	fmt.Println(slice, "tamanho do slice: ", len(slice))
}
