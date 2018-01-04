package main

import (
	"github.com/user/funcoes_basico/matematica"
	"fmt"
)

func main(){
	resultado := matematica.Calculo(matematica.Multiplicador,2,4)
	fmt.Printf("o resultado da multiplicação é: %d\r\n", resultado)
	resultado = matematica.Soma(2,3)
	fmt.Printf("o resultado da soma é: %d\r\n", resultado)
}
