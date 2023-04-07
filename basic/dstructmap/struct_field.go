package dstructmap

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Interval struct {
	start int
	end   int
}

func Gorun123412() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Printf("%v", ms)

	// new(Type) 和 &Type{}是等价的
	ms2 := new(struct1)
	ms2.i1 = 12
	fmt.Println(ms2)

	ms1 := &struct1{12, 12, "fsdfs"}
	fmt.Println(ms1)

	//intr := Interval{0, 3}
	//intr := Interval{end: 5, start: 1}
	//intr := Interval{end: 5}

	intr := Interval{0, 2}
	fmt.Println(intr)
	fmt.Println(intr.end, intr.start)

	intr1 := Interval{end: 13, start: 1254}
	fmt.Println(intr1)
	fmt.Println(intr1.end, intr1.start)

	intr2 := Interval{end: 4}
	fmt.Println(intr2)
	fmt.Println(intr2.end, intr2.start)

}
