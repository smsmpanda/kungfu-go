package dstructmap

import (
	"fmt"
	"reflect"
)

type TagType struct {
	name string "this is name"
	age  int    "this is age"
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func Gorun123445() {
	st := TagType{name: "Mo", age: 32}

	for i := 0; i < 2; i++ {
		refTag(st, i)
	}

	fmt.Printf("I'm %s %d years old", st.name, st.age)

	outer := new(outerS)
	outer.b = 1
	outer.c = 123.3
	outer.int = 21
	outer.innerS.in1 = 12
	outer.innerS.in2 = 13

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	outer2 := outerS{1, 123.1, 3, innerS{in1: 12, in2: 324}}
	fmt.Println("outer2 is:", outer2)
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
