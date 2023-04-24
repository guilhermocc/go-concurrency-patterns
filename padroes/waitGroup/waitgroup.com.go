package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementa o contador ao finalizar a goroutine
	fmt.Printf("Worker %d começou\n", id)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Printf("Worker %d terminou\n", id)
}

func main() {
	var wg sync.WaitGroup // cria o waitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1) // incrementa o contador para cada goroutine adicionada
		go worker(i, &wg)
	}
	wg.Wait() // espera todas as goroutines finalizarem
	fmt.Printf("Todos os workers terminaram\n")
}

//func worker(id int) {
//	fmt.Printf("Worker %d começou\n", id)
//	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
//	fmt.Printf("Worker %d terminou\n", id)
//}

//func main() {
//	for i := 1; i <= 5; i++ {
//		go worker(i)
//	}
//
//	time.Sleep(1 * time.Second)
//	fmt.Printf("Todos os workers talvez tenham terminado\n")
//}
