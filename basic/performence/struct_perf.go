package performence

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func (p *Person) String() string {
	return "the current people's name is " + p.name + " and him age is " + strconv.FormatInt(int64(p.age), 16)
}

func GorunStruct() {
	p := new(Person)
	p.age = 12
	p.name = "jiangrui"
	fmt.Printf("%v \n", p)

	p1 := &Person{"shilijuan", 12}
	fmt.Println(p1.name)

	p1.String()
}
