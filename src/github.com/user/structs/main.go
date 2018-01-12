package main

import (
	"encoding/json"
	"github.com/user/structs/model"
	"fmt"
)


func main(){
	casa := model.Imovel{} //não é uma expressão, precisa de "{}"
	casa.Nome = "casa azul"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("A casa é: %+v\r\n",casa) //%+v printa os nomes E o valor dos campos
	fmt.Printf("O valor da casa é: %d\r\n",casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}

/*
script da aula de struct básica

//Imovel é um tipo de dados
type Imovel struct{
	X int
	Y int
	Nome string
	valor int			//exporta ou não os campos da struct
}

func main(){
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n",casa) //%+v printa os nomes E o valor dos campos

	apartamento := Imovel{17,56,"meu ap",1000001}
	fmt.Printf("O apartamento é: %+v\r\n",apartamento)

	chacara := Imovel{
		Y:85,
		Nome: "Fontinha",
		X:22,
		valor: 55000005,
	}
	fmt.Printf("A chácara é: %+v\r\n",chacara)

	casa.Nome = "Lar de crianças peculiares"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 21
	fmt.Printf("A nova casa é: %+v\r\n",casa) //%+v printa os nomes E o valor dos campos
}
*/	