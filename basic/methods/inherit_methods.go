package methods

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person2 struct {
	base       *Base
	FirstName1 string
	LastName1  string
}

type Employee struct {
	Person2
	salary float32
}

func Gorun41241() {
	idjb := new(Base)
	idjb.id = "007"
	jb := Person2{idjb, "James", "Bond"}
	e := &Employee{jb, 100000.}
	fmt.Printf("ID of our hero: %v\n", e.base.Id())

	e.base.SetId("007B")
	fmt.Printf("The new ID of our hero: %v\n", e.base.Id())
}
