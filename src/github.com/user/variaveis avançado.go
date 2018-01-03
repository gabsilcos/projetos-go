package main

import (
	"fmt"
)

var (
	endereco, telefone string
	quantidade         int
	comprou            bool
	valor              float64
	palavras           rune //sem maiúscula é privada, com maiúscula é pública! públicas precisam de comentários
	//e declarações individuais

	//Saidinha é bem saidinha e gosta de passear lá fora xD
	Saidinha rune
)

func main() {
	teste := "valor de teste"
	fmt.Printf("endereço: %s\r\n", endereco)
	fmt.Printf("telefone: %s\r\n", telefone)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("teste! arrrrg! - %v\r\n", teste)
	fmt.Printf("palavras rúnicas arrrrg! - %v\r\n", palavras)
}
