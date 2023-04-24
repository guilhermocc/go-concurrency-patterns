package main

import (
	"fmt"
	"sync"
)

type Pessoa struct {
	Nome  string
	Idade int
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Pessoa)
	},
}

func main() {
	p := pool.Get().(*Pessoa)
	p.Nome = "João"
	p.Idade = 30
	fmt.Printf("Pessoa antes de devolver ao pool: %+v\n", *p)
	pool.Put(p)
	p2 := pool.Get().(*Pessoa)
	fmt.Printf("Pessoa retirada do pool: %+v\n", *p2)
}
