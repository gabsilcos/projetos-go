package main

import (
	"fmt"

	"github.com/user/structs/model"
)

var cache map[string]model.Imovel //map[tipo de dado que vai usar]tipo de dado que vai armazenar
//mapas são iniciados com o valor null, os únicos

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.X = 18
	casa.Y = 25
	casa.Nome = "Casa molhada"
	casa.SetValor(120000)

	cache["teste de casa molhada"] = casa

	fmt.Println("Há uma casa molhada no cache?")
	imovel, achou := cache["teste de casa molhada"]
	if achou {
		fmt.Printf("Sim, achei: %+v\r\n", imovel)
	}

	casa2 := model.Imovel{}
	casa2.X = 17
	casa2.Y = 23
	casa2.Nome = "AP seco"
	casa2.SetValor(150000)

	cache[casa2.Nome] = casa2

	fmt.Println("Quantos elementos há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave: [%s] = %+v\r\n", chave, imovel)
	}

	_, achou = cache["teste de casa molhada"]
	if achou {
		//delete(cache, imovel.Nome)
		delete(cache, "teste de casa molhada")
		fmt.Println("Quantos elementos há no cache? ", len(cache))
		for chave, imovel := range cache {
			fmt.Printf("Chave: [%s] = %+v\r\n", chave, imovel)
		}
	}
}
