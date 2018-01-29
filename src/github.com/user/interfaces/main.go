package main

import (
	"fmt"

	"github.com/user/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo not so evil"

	QueroAcordarComUmCacarejo(jojo)
	Duck(jojo)
}

//QueroAcordarComUmCacarejo são desejos bizarros do instrutor
func QueroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

//Duck são desejos bizarros do instrutor
func Duck(p model.Pato) {
	fmt.Println(p.Grasna())
}
