package main

import (
	"context"
	"fmt"
	"time"
)

//func main() {
//	for i := 0; i < 10; i++ {
//		go func(n int) {
//			ticker := time.NewTicker(time.Second)
//			for {
//				<-ticker.C
//				fmt.Println("Goroutine ", n)
//			}
//		}(i)
//	}
//
//	for {
//		// Programa principal rodando...
//	}
//}

func main() {
	ctx := context.Background()

	contextoComTimer, cancelFn := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFn()

	for i := 0; i < 10; i++ {
		go func(ctx context.Context, n int) {
			ticker := time.NewTicker(time.Second)
			for {
				select {
				case <-contextoComTimer.Done():
					fmt.Println("Goroutine ", n, " finalizada")
					return
				case <-ticker.C:
					fmt.Println("Goroutine ", n)
				}
			}
		}(contextoComTimer, i)
	}

	for {
		// Programa principal rodando...
	}
}
