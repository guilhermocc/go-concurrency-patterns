package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, eg errgroup.Group) error {
	defer fmt.Printf("Worker %d começou\n", id)
	valorParaSleep := rand.Intn(2000)
	if valorParaSleep > 1000 {
		return fmt.Errorf("Erro worker %d, tempo de espera muito alto %d", id, valorParaSleep)
	}
	time.Sleep(time.Duration(valorParaSleep) * time.Millisecond)
	fmt.Printf("Worker %d terminou\n", id)
	return nil
}

func main() {
	// cria error group eg
	var eg errgroup.Group // cria o errgroup
	for i := 1; i <= 5; i++ {
		eg.Go(func() error {
			return worker(i, eg)
		})
	}

	err := eg.Wait() // espera todas as goroutines finalizarem
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Printf("Todos os workers terminaram com sucesso\n")
	}

}
