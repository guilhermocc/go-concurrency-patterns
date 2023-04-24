package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}

//func main() {
//	counter := make(chan int)
//	defer close(counter)
//
//	go func() {
//		counter <- 0
//	}()
//
//	for i := 0; i < 1000; i++ {
//		go func(counter chan int) {
//			value := <-counter
//			counter <- value + 1
//		}(counter)
//	}
//
//	time.Sleep(1 * time.Second)
//	fmt.Println(<-counter)
//}

//func main() {
//	counter := 0
//	mtx := sync.Mutex{}
//
//	for i := 0; i < 1000; i++ {
//		go func() {
//			mtx.Lock()
//			counter++
//			mtx.Unlock()
//		}()
//	}
//
//	time.Sleep(1 * time.Second)
//	fmt.Println(counter)
//}
