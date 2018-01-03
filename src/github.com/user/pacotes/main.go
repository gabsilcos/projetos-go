package main

import (
	"fmt"
 	"github.com/user/pacotes/prefixo"
 	"github.com/user/pacotes/operadora"
)
//NomeDoUsuario é Nome do usuário do Sistema
var NomeDoUsuario = "BB"

func main() {
	fmt.Printf("Nome do Usuário: %v\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo com nome da Operadora: %s\r\n", operadora.PrefixoDaCapitalDaOperadora)
}
