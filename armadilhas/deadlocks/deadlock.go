package main

import (
	"fmt"
	"sync"
)

//func main() {
//	ch := make(chan int)
//	<-ch
//	go func() {
//		ch <- 1
//	}()
//}

//func main() {
//	var mu sync.Mutex
//	go func() {
//		mu.Lock()
//		fmt.Println("goroutine 1 adquiriu o Mutex")
//	}()
//	time.Sleep(1 * time.Second)
//	mu.Lock()
//	fmt.Println("goroutine 2 adquiriu o Mutex")
//	mu.Unlock()
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("goroutine")
	}()
	wg.Wait()
}
