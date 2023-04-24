package main

import "fmt"

func main() {
	// Gerador de dados
	numeros := geraNumeros(10)

	// Estágios de processamento
	pipeline := quadradoNumeros(dobraNumeros(numeros))

	// Consumidor final
	for num := range pipeline {
		fmt.Println(num)
	}
}

func geraNumeros(quantidade int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for numero := 0; numero < quantidade; numero++ {
			out <- numero
		}
	}()
	return out
}

func dobraNumeros(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for numero := range in {
			out <- numero * 2
		}
	}()
	return out
}

func quadradoNumeros(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}
