//Package model pacote contendo structs
package model

import (
	"errors"
)

//Imovel é um tipo de dados
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int    //exporta ou não os campos da struct, valor não exportado não recebe tag
}

//ErrValorDeImovelInvalido erro personalizado
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

//ErrValorDeImovelMuitoAlto erro personalizado, apresentado quanto SetValor é acima da faixa permitida
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto para o imóvel")

//SetValor é a alteração seletiva do valor de Imovel.valor
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
	}
	i.valor = valor
	return
}

//GetValor é a leitura seletiva do valor de Imovel.valor
func (i *Imovel) GetValor() int {
	return i.valor
}
