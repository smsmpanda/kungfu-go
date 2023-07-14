package dataformatter

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

var content string
var vc VCard

func GorunGob2() {

	file, _ := os.Open("resources/vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc.Addresses)
	fmt.Println(vc.FirstName)
	fmt.Println(vc.LastName)
	fmt.Println(vc.Remark)
}
