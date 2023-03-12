package examples

import "fmt"

func Going123141() {
	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}

	items := [...]int{10, 201, 301, 410}
	for _, item := range items {
		item *= 2
	}

	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum(a))

	slice2 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice2); i++ {
		slice2 = slice2[0 : i+1]
		slice2[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice2))
	}
}

func Sum(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
