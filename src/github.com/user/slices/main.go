package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums)) //cap é a capacidade do array/slice
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums)) //cap é a capacidade do array/slice
	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	cidades[0] = "c1"
	cidades[1] = "c2"
	cidades[2] = "c3"
	cidades[3] = "c4"
	cidades[4] = "c5"
	fmt.Println(cidades, len(cidades), cap(cidades))
	capitais[1] = "Machupichu"
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\r\n", indice, cidade)
	}

	capitaisAsia := cidades[3:5] /*
		"o primeiro item começa com 0
		e o segundo item começa com 1" xD*/
	fmt.Println("Capitais a partir do 3° item: ", capitaisAsia, len(capitaisAsia), cap(capitaisAsia))
	temp1 := cidades[:2]
	fmt.Println("Capitais até o 2° item: ", temp1, len(temp1), cap(temp1))
	temp2 := cidades[len(cidades)-2:]
	fmt.Println("Dois últimos itens: ", temp2, len(temp2), cap(temp2))

	//remover slice do array: porco
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	fmt.Println("Copia os itens anteriores: ", temp)
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	fmt.Println("Pula um e adiciona o resto ao final: ", temp)
	for indice, cidade := range temp {
		fmt.Printf("Cidade[%d] = %s\r\n", indice, cidade)
	}
	//copia de volta pro original
	copy(cidades, temp)
	fmt.Println(cidades)
	//temp é menor que cidades e o último elemento de cidades não é sobrescrito
}
