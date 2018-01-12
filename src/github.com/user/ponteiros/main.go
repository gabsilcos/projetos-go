package main

import (
	"fmt"
)

//Imovel é um tipo de dados
type Imovel struct{
	X int
	Y int
	Nome string
	valor int			//exporta ou não os campos da struct
}

func main(){
	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17,28,"chacara e só",280000}
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
	//casa.Nome = "Lar de crianças peculiares"
	//casa.valor = 450000
	//casa.X = 18
	//casa.Y = 21
	//fmt.Printf("A nova casa é: %+v\r\n",casa) //%+v printa os nomes E o valor dos campos
}

func mudaImovel(imovel *Imovel){
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}