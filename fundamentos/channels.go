package main

import (
	"fmt"
)

// produtor recebe um channel "send-only" (somente envio) e envia dados para o channel.
func produtor(channel chan<- int) {
	for i := 1; i <= 5; i++ {
		// A seta indica o sentido do fluxo de dados, neste caso estamos enviando o dado "i" para o channel
		channel <- i
	}
	// Precisamos fechar o channel, pois o consumidor depende do fechamento do channel para parar.
	close(channel)
}

// consumidor recebe um channel "receive-only" (somente recebimento) e recebe dados do channel.
func consumidor(channel <-chan int) {
	for {
		// A seta indica o sentido do fluxo de dados, neste caso estamos recebendo o dado do channel, e
		// armazenando na variável valor
		valor, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(valor)
	}
}

//func main() {
//	channel := make(chan int)
//	go produtor(channel)
//	consumidor(channel)
//}
