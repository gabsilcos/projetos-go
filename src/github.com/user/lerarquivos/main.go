package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	//scanner lê linha-a-alinha
	//scanner := bufio.NewScanner(arquivo)
	//for scanner.Scan() {
	//	linha := scanner.Text()
	//		fmt.Println("O conteúdo da linha é: ", linha)
	//	}

	//leitor csv é específico para tratamento de arquivos .csv
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
		return
	}
	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indiceItem, item)
		}
	}
	arquivo.Close()

}
