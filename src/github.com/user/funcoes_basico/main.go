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
	resultado = matematica.Calculo(matematica.Divisor,8,4)
	fmt.Printf("o resultado da divisão é: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(9,4)
	fmt.Printf("o resultado da divisão é: %d e o resto é %d\r\n", resultado,resto)
}
