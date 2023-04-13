package methods

import (
	"fmt"
)

type TwoInts struct {
	a int
	b int
}

type IntTypeAraay []int

func (ita IntTypeAraay) Sum() (s int) {
	for _, v := range ita {
		s += v
	}
	return
}

func (b *B) Change() {
	b.n = 123
}

func (b B) Change1() {
	b.n = 200
}

func Gorun4123() {
	fmt.Println("method ....")

	twoint := TwoInts{a: 12, b: 144}
	fmt.Printf("the sum is %d\n", twoint.AddThem())

	fmt.Printf("the parameter sum is %d\n", twoint.AddToParam(123))

	ita := IntTypeAraay{1, 2, 3, 4}
	fmt.Println(ita.Sum())

}

func (twoInt *TwoInts) AddThem() int {
	return twoInt.a + twoInt.b
}

func (twoInt *TwoInts) AddToParam(param int) int {
	return twoInt.a + twoInt.b + param
}
