package main

import (
	"fmt"
)

func main(){

	for i := 0; i < 10; i++{
		//de zero a 9
		fmt.Printf("i vale %d\r\n", i)
	}

	valor := 0
	teste := true

	for teste{
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O número é ", valor)
	}

	for {
		valor--
		fmt.Println("O número é ", valor)
		if valor == 0 {break}
	}

	texto := "Eu adoro dormir bastante, direto e até bem tarde :D"

	for indice, letra := range texto{
		fmt.Printf("Texto[%d] = %q\r\n", indice, letra)
	}
}