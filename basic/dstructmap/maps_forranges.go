package dstructmap

import "fmt"

func Gorun412() {
	fmt.Println("for map range...")

	items := make([]map[int]int, 10)
	for i := range items {
		items[i] = make(map[int]int, 2)
		items[i][1] = 2 * i
		items[i][2] = 2 * i * i
	}

	fmt.Printf("Version A: Value of items: %v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}

	fmt.Printf("Version B: Value of items: %v\n", items2)
}
