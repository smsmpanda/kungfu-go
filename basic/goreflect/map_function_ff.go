package goreflect

import "fmt"

type obj interface{}

func Gorun13131() {
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	is1 := []obj{0, 1, 2, 3}
	res1 := mapFunc(mf, is1)
	for _, v := range res1 {
		fmt.Println(v)
	}

	println()

	ssl := []obj{"0", "1", "2", "3"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}
