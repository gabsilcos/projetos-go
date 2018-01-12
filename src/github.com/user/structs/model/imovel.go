//Package model pacote contendo structs
package model



//Imovel é um tipo de dados
type Imovel struct{
	X		int		`json:"coordenada_x"`
	Y		int		`json:"coordenada_y"`
	Nome	string	`json:"nome"`
	valor	int		//exporta ou não os campos da struct, valor não exportado não recebe tag
}

//SetValor é a alteração seletiva do valor de Imovel.valor
func (i *Imovel) SetValor (valor int){
	i.valor = valor
}

//GetValor é a leitura seletiva do valor de Imovel.valor
func (i *Imovel) GetValor() int {
	return i.valor
}