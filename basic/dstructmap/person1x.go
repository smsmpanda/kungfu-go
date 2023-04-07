package dstructmap

import (
	"fmt"
	"strings"
)

type Personx struct {
	firstName string
	lastName  string
}

func upperPerson(p Personx) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func Gorunfsdfsdf() {
	// 1- struct as a value type:
	var pers1 Personx
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upperPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2 - struct as a pointer:
	pers2 := new(Personx)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	upperPerson(*pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3 - struct as a literal:
	pers3 := &Personx{"Chris", "Woodward"}
	upperPerson(*pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
