package main

import (
	"fmt"
)

func main(){
	meses := 0
	situacao := true
	cidade := "terra de nenhuma terra"

	//< > != == <= >= && ||
	if meses <= 6 && meses > 0 {
		fmt.Println("Esse credor deve há pouco tempo")
	}else if meses == 0 {
		fmt.Println("Não deve")
	}else{
		fmt.Println("Condenado")
	}

	if situacao { //testa true por default
		fmt.Println("Devedor")
	}
	if !situacao { //testa true por default
		fmt.Println("Adimplente")
	}

	if cidade != "terra de nenhuma terra" {//case sensitive
		fmt.Println("Talvez não more tão mal")
	}else{
		fmt.Println("Mora mal")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status{
		//as variáveis passadas aqui só são válidas dentro do if
		fmt.Println("Qual a situação do cliente?", descricao)
	}else{
		fmt.Println("Qual a situação do cliente?", descricao)
	}

	fmt.Println("Fim da consulta.")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool){
	if meses > 0 {
		status = true
		descricao = "Devendo, feladaputa"

		return
	}
	descricao = "Não-devedor, abençoado"

	return
}