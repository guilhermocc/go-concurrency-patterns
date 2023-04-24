package main

import (
	"sync"
	"time"
)

func fanOut(numeros []int, numChannels int) []<-chan int {
	if numChannels > len(numeros) {
		numChannels = len(numeros)
	}

	valorPorCanal := len(numeros) / numChannels

	var channels []<-chan int
	for i := 0; i < numChannels; i++ {
		channel := make(chan int)
		go func(valores []int, canal chan<- int) {
			for _, valor := range valores {
				time.Sleep(1 * time.Second)
				canal <- valor
			}
			close(canal)
		}(numeros[i*valorPorCanal:(i+1)*valorPorCanal], channel)
		channels = append(channels, channel)
	}
	return channels
}

func main() {
	canais := fanOut([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	wg := sync.WaitGroup{}
	for _, canal := range canais {
		wg.Add(1)
		go func(canal <-chan int) {
			defer wg.Done()
			for numero := range canal {
				println(numero)
			}
		}(canal)
	}
	wg.Wait()
}
