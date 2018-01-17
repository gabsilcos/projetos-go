package main

//go build -o appDaAula
//GOOS=windows GOARCH=386 go build -o appWindows.exe
//GOOS=linux GOARCH=386 go build -o appLinux
//GOOS=darwin GOARCH=amd64 go build -o hello_osx_amd64 hello.go
//-v github.com/user/structs

import (
	"encoding/json"
	"fmt"

	"github.com/user/structs/model"
)

func main() {
	casa := model.Imovel{} //não é uma expressão, precisa de "{}"
	casa.Nome = "casa azul"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(100000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição do valor da casa:", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Vai roubar a mãe!")
		}
		return
	}

	fmt.Printf("A casa é: %+v\r\n", casa) //%+v printa os nomes E o valor dos campos
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)
	//o erro é um ponteiro que aponta pra um ojeto que implementa a interface 'error'
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do obj JSON", err.Error())
		//existe o panic pra tratamento de erros, que mata tudo
		return
	}
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
