package examples

import (
	"fmt"
)

func Going1() {
	x := mymin(1, 3, 2, 24)
	fmt.Printf("The minimum is: %d\n", x)

	slice := []int{15, 35, 13, 125, 63}
	x = mymin(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func mymin(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	for i, n := range s {
		fmt.Printf(" %d - %d\n", i, n)
	}
	return min
}
