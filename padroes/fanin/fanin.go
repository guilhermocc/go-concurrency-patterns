package main

import (
	"sync"
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

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(channel <-chan int) {
			defer wg.Done()
			for n := range channel {
				out <- n
			}
		}(channel)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	numerosAlice := geraNumeros(1, 5)
	numerosBob := geraNumeros(6, 10)

	numeros := fanIn(numerosAlice, numerosBob)

	for numero := range numeros {
		println("Numero: ", numero)
	}
}
