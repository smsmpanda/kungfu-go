package examples

import "fmt"

func Going123123() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	sl4 := []int{1, 2, 3}
	appendByte(sl4, 10, 20, 30, 40, 50)
}

func appendByte(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
