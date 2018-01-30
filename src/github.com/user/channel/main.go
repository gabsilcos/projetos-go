package main

//a rtoca de dados por canais trava o processamento até que a função da vez termmine de trabalhar com o canal
//semelhante ao que acontece no dataflow do labview

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	for {
		canal <- "ping" //canal recebe
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong" //canal recebe
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal //canal entrega
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var canal chan string //objeto canal transfere dados entre funções que recebem canais, pode ter tipos
	canal = make(chan string)
	go pinger(canal)
	go ponger(canal)
	go impressora(canal)
	//aqui a execução "não é sequencial": uma vez iniciados assincronamente, serão executados indefinidamente
	//como o pinger é executado primeiro, entrega dados primeiro

	var entrada string
	fmt.Scanln(&entrada) //lê entradas no teclado
	//depois que o scan é executado (e responde) nada é posto no canal, e, sem loop, finda o programa

}
