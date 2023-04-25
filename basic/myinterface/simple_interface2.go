package myinterface

import (
	"fmt"
	"strconv"
)

type Simpler1 interface {
	Get() int
	Set(int)
}

type Simple1 struct {
	i int
}

func (p *Simple1) Get() int {
	return p.i
}

func (p *Simple1) Set(u int) {
	p.i = u
}

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Set(u int) {
	p.j = u
}

func fI2(it Simpler1) int {
	switch it.(type) {
	case *Simple1:
		it.Set(5)
	case *RSimple:
		it.Set(5)
	default:
		return 99
	}
	return 0
}

type IM interface {
	Get() int
	Set(int)
	String() string
}

type IMer struct {
	id int
}

func (im *IMer) Get() int {
	return im.id
}

func (im *IMer) Set(id int) {
	im.id = id
}

func (im *IMer) String() string {
	return "I'm a imer and id is " + strconv.Itoa(im.id)
}

func Gorun14124() {
	var s Simple1
	fmt.Println(fI2(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple
	fmt.Println(fI2(&r))

	im := new(IMer)
	im.Set(12)

	var iim IM = im
	id := im.Get()
	fmt.Println(id)
	fmt.Println(im)

	if v, ok := iim.(IM); ok {
		fmt.Printf("im implements String(): %s\n", v.String())
		fmt.Println(v)
	}

	switch t := iim.(type) {
	case *IMer:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
