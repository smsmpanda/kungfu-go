package methods

import (
	"fmt"
	"strconv"
)

type Stack struct {
	ix   int
	data []int
}

func Gorun() {
	st := new(Stack)
	st.Push(2)
	fmt.Printf("%v\n", st)
	st.Push(4)
	fmt.Printf("%v\n", st)
	st.Push(6)
	fmt.Printf("%v\n", st)
	st.Push(8)
	fmt.Printf("%v\n", st)
	st.Push(10)
	fmt.Printf("%v\n", st)

	st.Pop()
	fmt.Printf("%v\n", st)
	st.Pop()
	fmt.Printf("%v\n", st)
	st.Pop()
	fmt.Printf("%v\n", st)
	st.Pop()
	fmt.Printf("%v\n", st)
	st.Pop()
	fmt.Printf("%v\n", st)
}

func (st *Stack) Push(n int) {
	st.data = append(st.data, n)
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix <= 0 {
		return 0
	}

	v := st.data[st.ix-1]
	st.data = st.data[:st.ix-1]
	st.ix--
	return v
}

func (st *Stack) String() string {
	str := ""
	for ix, v := range st.data {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}
