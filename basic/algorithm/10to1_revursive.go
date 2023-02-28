package algorithm

import "fmt"

func Algorithmgo(n int) {
	// if n >= 1 {
	// 	fmt.Println(n)
	// 	n--
	// 	Algorithmgo(n)
	// }
	// printrec(1)
	fmt.Println(jianchen(15))
}

func printrec(i int) {
	if i > 3 {
		return

	}
	printrec(i + 1)
	fmt.Printf("%d ", i)
}

func jianchen(n int) int {
	if n <= 1 {
		return 1
	} else {
		fmt.Printf("%d ", n)
		return n * jianchen(n-1)
	}
}
