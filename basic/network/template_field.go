package network

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

func Gorunend() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := Person{Name: "Mary", nonExportedAgeField: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
