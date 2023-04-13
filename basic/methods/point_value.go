package methods

import (
	"fmt"
)

type B struct {
	n int
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func Gorun512() {
	b := B{n: 12}
	fmt.Printf("prev %d \n", b.n)
	b.Change()
	fmt.Printf("after %d \n", b.n)
	b.Change1()
	fmt.Printf("value paramter the result is %d \n", b.n)

	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
}
