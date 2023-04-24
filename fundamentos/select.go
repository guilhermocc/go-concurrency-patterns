package main

import "time"

func contaSemSelect() {
	numerosAlice := make(chan int)
	numerosBob := make(chan int)

	go func(channel chan<- int) {
		for i := 1; i <= 5; i++ {
			channel <- i
			time.Sleep(1 * time.Second)
		}
		close(channel)
	}(numerosAlice)

	go func(channel chan<- int) {
		for i := 6; i <= 10; i++ {
			channel <- i
			time.Sleep(1 * time.Second)
		}
		close(channel)
	}(numerosBob)

	for i := range numerosAlice {
		println("Alice: ", i)
	}
	for i := range numerosBob {
		println("Bob: ", i)
	}
}

func contaComSelect() {
	numerosAlice := make(chan int)
	numerosBob := make(chan int)

	go func(channel chan<- int) {
		for i := 1; i <= 5; i++ {
			channel <- i
			time.Sleep(1 * time.Second)
		}
	}(numerosAlice)

	go func(channel chan<- int) {
		for i := 6; i <= 10; i++ {
			channel <- i
			time.Sleep(1 * time.Second)
		}
	}(numerosBob)

	for i := 1; i <= 10; i++ {
		select {
		case numero := <-numerosAlice:
			println("Alice: ", numero)
		case numero := <-numerosBob:
			println("Bob: ", numero)
		}
	}
}

func main() {
	now := time.Now()
	contaComSelect()
	println("Tempo: ", time.Since(now).Seconds())
}
