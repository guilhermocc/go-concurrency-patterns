package main

import "time"

func conta(nome string, inicio int, fim int) {
	for i := inicio; i <= fim; i++ {
		println(nome+": ", i)
		time.Sleep(1 * time.Second)
	}
}

//func main() {
//	go conta("Alice", 1, 5)
//	go conta("Bob", 5, 10)
//
//	time.Sleep(5 * time.Second)
//}

//func main() {
//	conta("Alice", 1, 5)
//	conta("Bob", 5, 10)
//}

//func main() {
//	go conta("Alice", 1, 5)
//	go conta("Bob", 5, 10)
//}
