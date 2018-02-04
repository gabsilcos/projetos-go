package main

import (
	"fmt"
	"net/http"

	"github.com/user/servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!") //F print recebe uma fç
	})
	http.HandleFunc("/funcao", manipulador.Funcao) //não é necessário, várias rotas com seus respectivos
	http.HandleFunc("/ola", manipulador.Ola)       //                  manipuladores definidos
	fmt.Println("Estou escutando, comandante...")  //não é necessário
	http.ListenAndServe(":8181", nil)              //é necessário: 'nil' pq o handler já foi criado acima
}
