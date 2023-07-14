package dataformatter

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func Gorunxml() {
	fmt.Println("xml")
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	fmt.Println(p)
}
