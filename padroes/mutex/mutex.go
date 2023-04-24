package main

import (
	"fmt"
	"sync"
	"time"
)

type ContadorSeguro struct {
	valores    map[string]int
	valoresMtx sync.Mutex
}

func (c *ContadorSeguro) Inc(key string) {
	c.valoresMtx.Lock() // bloqueia o acesso ao mapa
	c.valores[key]++
	c.valoresMtx.Unlock() // desbloqueia o acesso ao mapa
}

func (c *ContadorSeguro) Value(key string) int {
	c.valoresMtx.Lock()         // bloqueia o acesso ao mapa
	defer c.valoresMtx.Unlock() // desbloqueia o acesso ao mapa após o return
	return c.valores[key]
}

//func main() {
//	c := ContadorSeguro{valores: make(map[string]int)}
//	for i := 0; i < 1000; i++ {
//		go c.Inc("algumaChave")
//		go fmt.Println(c.Value("algumaChave"))
//	}
//	time.Sleep(time.Second)
//	fmt.Println(c.Value("algumaChave"))
//}

type Contador struct {
	valores map[string]int
}

func (c *Contador) Inc(chave string) {
	c.valores[chave]++
}

func (c *Contador) Valor(chave string) int {
	return c.valores[chave]
}

func main() {
	c := Contador{valores: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("algumaChave")
		go fmt.Println(c.Valor("algumaChave"))
	}
	time.Sleep(time.Second)
	fmt.Println(c.Valor("somekey"))
}
