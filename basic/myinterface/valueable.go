package myinterface

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func Gorun4124124() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	var r io.Reader
	r = os.Stdin // see 12.1
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
}
