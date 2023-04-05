package dstructmap

import (
	"container/list"
	"fmt"
	"sort"
	"unsafe"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func Gorun41234() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")

	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

	fmt.Println("0000000000000000000")

	a := []string{"g", "a"}
	sort.Strings(a)
	fmt.Printf("%v\n", a)

	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)
	// fmt.Println("Here is the double linked list:\n", lst)
	for e := lst.Front(); e != nil; e = e.Next() {
		// fmt.Println(e)
		fmt.Println(e.Value)
	}

	var ii int64 = 10
	size := unsafe.Sizeof(ii)
	fmt.Println("The size of an int is: ", size)
}
