package algorithm

import "fmt"

func Algorithmgo1() {
	fmt.Println(fibonacci(4))
}

func fibonacci(n int) (res int, i int) {
	if n <= 1 {
		res = 1
	} else {
		prev, _ := fibonacci(n - 1)
		next, _ := fibonacci(n - 2)
		res = prev + next
	}
	i = n
	return
}

func fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci1(n-1) + fibonacci1(n-1)
	}
	return
}
