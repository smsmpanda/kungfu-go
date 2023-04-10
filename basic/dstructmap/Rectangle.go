package dstructmap

import (
	"fmt"
	"unsafe"
)

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

type File struct {
	Fd   int
	Name string
}

type ExpStruct struct {
	Mi1 int
	Mf1 float32
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func Gorun12312() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area())
	fmt.Println("Rectangle perimeter is: ", r1.Perimeter())

	file := NewFile(12, "pxixixi")
	fmt.Println(file.Fd, file.Name)

	size := unsafe.Sizeof(file)
	fmt.Print(size)
}
