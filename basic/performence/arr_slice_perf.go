package performence

import (
	"fmt"
	"strings"
)

func GorunMapSliceArray() {
	// arr1 := new([3]int)
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 3
	// for _, v := range arr1 {
	// 	fmt.Println(v)
	// }

	// slice := make([]int, 31)
	// slice[0] = 1
	// slice[10] = 1
	// slice[20] = 1
	// slice[30] = 1

	arr1 := []string{"s", "k", "i", "l"}
	fmt.Println(strings.Join(arr1, "-"))

	seaon := make(map[string]int)
	seaon["spring"] = 2
	fmt.Println(seaon["spring"])

	if v, ok := seaon["summer"]; ok {
		fmt.Println("found", v)
	} else {
		fmt.Println("not found")
	}
}
