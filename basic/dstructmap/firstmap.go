package dstructmap

import (
	"fmt"
)

// 1、key可以是任意可以用== != 操作符比较的类型，比如string int float
// 2、数组 切片 结构体是不可以的 若要用结构体作为key可以提供Key() Hash()方法，这样可以通过结构体
// 的域计算出唯一的数字或者字符串的key
// 3、指针 接口时可以作为key的
func Gorun1() {
	var map1 map[string]string
	fmt.Println(map1)

	var maplist map[string]int
	var mapAssigned map[string]int

	maplist = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = maplist

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	if _, ok := maplist["two"]; ok {
		fmt.Println("find key")
	}
}
