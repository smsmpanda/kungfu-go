package goreflect

import (
	"fmt"
	"io"
)

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
}

func (bird *Bird) Quack() {
	fmt.Println("I'm quacking")
}

func (bird *Bird) Walk() {
	fmt.Println("I'm walking")
}

type xmlWriter interface {
	WriteXML(w io.Writer) error
}

func Gorun324ff() {
	b := new(Bird)
	DuckDance(b)
}
