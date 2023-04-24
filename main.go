package main

import "concurrency-patterns-go/padroes/generator"

func main() {
	// Recebendo dados do channel, neste exemplo estamos usando o range para receber os dados do channel.
	for numero := range generator.Numeros() {
		println("Numero: ", numero)
	}
}
