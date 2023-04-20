package myinterface

import "fmt"

type Simpler interface {
	Get() int
	Put(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Put(u int) {
	p.i = u
}

func fI(it Simpler) int {
	it.Put(5)
	return it.Get()
}

func Gorun() {
	var s Simple
	fmt.Println(fI(&s))
}
