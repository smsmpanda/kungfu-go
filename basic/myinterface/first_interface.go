package myinterface

import "fmt"

//类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

//实现某个接口的类型（除了实现接口方法外）可以有其他的方法。

//一个类型可以实现多个接口。

//接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func Gorun2131() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer

	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
