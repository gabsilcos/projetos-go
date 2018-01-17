package main

import (
	"fmt"
)

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1
	fmt.Println("Qual a capacidade deste array? ", len(teste), "itens")

	cantores := [2]string{"cavaleira", "derick"}
	fmt.Printf("Q q tem nessa budega? \n\r%v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Brasília", "Luanda", "Maputo"}
	fmt.Println("Qual a capacidade deste array? ", len(capitais), "itens")

	for indice, cidade := range capitais {
		fmt.Printf("A capital[%d] é %s \r\n", indice, cidade)
	}
}
