package myinterface

import (
	"fmt"
	"reflect"
)

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func Gorun13123123() {
	// v := new(Vector)
	// v.a = make([]Element, 10)
	// v.Set(0, 12)
	// v.Set(1, "hello world")

	// f := 3.32
	// fmt.Println(f)
	// t := reflect.TypeOf(f)
	// fmt.Println(t)
	// val := reflect.ValueOf(f)
	// fmt.Println(val)

	// var x float64 = 3.4
	// fmt.Println("type:", reflect.TypeOf(x))
	// v := reflect.ValueOf(x)
	// fmt.Println("value:", v)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind:", v.Kind())
	// fmt.Println("value:", v.Float())
	// fmt.Println(v.Interface())
	// fmt.Printf("value is %5.2e\n", v.Interface())
	// y := v.Interface().(float64)
	// fmt.Println(y)

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}
