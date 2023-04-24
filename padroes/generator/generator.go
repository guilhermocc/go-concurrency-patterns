package main

import (
	"time"
)

func geraNumeros(inicio, fim int) <-chan int {
	channel := make(chan int)
	go func() {
		for i := inicio; i <= fim; i++ {
			channel <- i
			time.Sleep(1 * time.Second)
		}
		close(channel)
	}()
	return channel
}

func main() {
	numerosAlice := geraNumeros(1, 5)
	numerosBob := geraNumeros(6, 10)

	for i := 1; i <= 10; i++ {
		select {
		case numero := <-numerosAlice:
			println("Alice: ", numero)
		case numero := <-numerosBob:
			println("Bob: ", numero)
		}
	}
}
