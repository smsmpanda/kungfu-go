package examples

import "fmt"

func Goin123g() {
	// var slice1 []int = make([]int, 10)
	// for i := 0; i < len(slice1); i++ {
	// 	slice1[i] = 5 * i
	// }

	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	// }
	// fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	// fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// var arr1 [6]int
	// var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i
	// }

	// // print the slice
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	// }

	// fmt.Printf("The length of arr1 is %d\n", len(arr1))
	// fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// // grow the slice
	// slice1 = slice1[0:4]
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	// }
	// fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))

	s = s[2:4]
	fmt.Println(len(s), cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s1)
	fmt.Println(s2)
}
