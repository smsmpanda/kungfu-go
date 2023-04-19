package methods

import (
	"fmt"
)

type Base1 struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()

	two1 := new(TwoInts2)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
	fmt.Printf("%v\n", two1)

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)

	fmt.Println(two1)
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func Gorun52352352() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

type TwoInts2 struct {
	a int
	b int
}

func (tn *TwoInts2) String() string {
	return "fsdfsdf"
}

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}
