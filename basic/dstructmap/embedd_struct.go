package dstructmap

import "fmt"

type A struct {
	ax, ay int
	bx     float32
}
type B struct {
	A
	bx, by float32
}

func Gorun1fd() {
	b := B{A{12, 12, 21.1}, 3.1, 12.3}
	fmt.Printf("%f", b.bx)
	fmt.Println(b)
}
