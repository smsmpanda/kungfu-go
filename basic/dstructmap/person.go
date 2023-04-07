package dstructmap

import (
	"fmt"
	"strings"
)

type Person struct {
	age       int
	lastName  string
	firstName string
}

type number struct {
	f float32
}

type nr number

func upPerson(p *Person) {
	p.lastName = strings.ToUpper(p.lastName)
}

func upPerson1(p Person) {
	p.lastName = "woshidaxiaongmao"
}

func Gorun2134124() {
	p := &Person{123, "jiang", "rui"}
	fmt.Printf("update before person's name is %s\n", p.lastName)
	upPerson(p)
	fmt.Printf("update after person's name is %s\n", p.lastName)

	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	fmt.Printf("The lastname is %s\n", pers1.lastName)
	upPerson1(pers1)
	fmt.Printf("The lastname is %s\n", pers1.lastName)

	a := number{5.0}
	b := nr{5.0}

	var c = number(b)
	fmt.Println(a, b, c)
}
