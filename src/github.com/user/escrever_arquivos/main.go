package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/user/escrever_arquivos/model"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
		return
	}

	//cria e abre o arquivo
	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON) //bufio.newWriter é um buffer de escrita
	escritor.WriteString("[\r\n")            //escreve strings diretamente
	for _, linha := range conteudo {         //"conteúdo é um objeto csv segmentado anteriormente"
		for indiceItem, item := range linha { //varre cada linha no conteudo
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade) //cria o json
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item ", item, ". Erro: ", err.Error())
			}
			escritor.WriteString("  " + string(cidadeJSON)) //escreve o json no arquivo
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush() //detona o buffer

	arquivoJSON.Close()
	arquivo.Close()
}
