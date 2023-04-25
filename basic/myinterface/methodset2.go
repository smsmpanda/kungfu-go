package myinterface

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func Gorun41242() {

	var lst List
	// VALID: Identical receiver type
	//CountInto(lst, 1, 10)

	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10) // VALID: Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}

//指针方法可以通过指针调用
//值方法可以通过值调用
//接收者是值的方法可以通过指针调用，因为指针首先被解引用
//接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
