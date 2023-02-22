package examples

import "fmt"

func Going() {
	a := []int{1, 5, 6, 7, 12, 61, 754, 12415}
	going(a...)
	going1(a)
}

func going(s ...int) {
	for i, n := range s {
		fmt.Printf("%d = %d \n", i, n)
	}
}

func going1(s []int) {
	for i, n := range s {
		fmt.Printf("%d = %d \n", i, n)
	}
}
