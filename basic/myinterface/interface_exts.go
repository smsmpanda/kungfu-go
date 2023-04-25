package myinterface

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type AreaInfoer interface {
	area() float32
}

type Triangle struct {
	di  float32
	gao float32
}

func (triangle Triangle) Area() float32 {
	return triangle.di * triangle.gao * 0.5
}

type Pers struct {
	firstName string
	lastName  string
}

func gI(any interface{}) int {
	if v, ok := any.(AreaInfoer); ok {
		return int(v.area())
	}
	return 0
}

type Perses []Pers

func (p Perses) Len() int { return len(p) }

func (p Perses) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}

func (p Perses) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Gorun14121234() {
	tri := Triangle{6, 4}
	fmt.Printf("the triangle's area is %.2f\n", tri.Area())
}
