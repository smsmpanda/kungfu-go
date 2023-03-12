package examples

import "fmt"

var s []int

func Going123333() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)

	filter(s, myfunc)

	s = append(s, 4)
}

func myfunc(data int) bool {
	if data%2 == 0 {
		return true
	}
	return false
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	s = ns
	return s
}

func filter(s []int, oushu func(int) bool) {
	for _, v := range s {
		if oushu(v) {
			fmt.Printf("%d is oushu!\n", v)
		}
	}
}
