package osrel

import (
	"fmt"
	"io/ioutil"
	"os"
)

func data(name string) string {
	f, _ := os.OpenFile(name, os.O_RDONLY, 0)
	defer f.Close() // idiomatic Go code!
	contents, _ := ioutil.ReadAll(f)
	return string(contents)
}

func Gorundefer() {
	var name string
	name = data("./resources/product.txt")
	fmt.Println(name)
}
