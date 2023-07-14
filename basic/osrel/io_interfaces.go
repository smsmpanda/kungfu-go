package osrel

import (
	"bufio"
	"fmt"
	"os"
)

func Gorunfsdfsfd() {
	fmt.Println("Hello World")

	fmt.Fprintln(os.Stdout, "Hello World")

	fmt.Fprintf(os.Stdout, "Hello World %s", "Jack")

	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "Hello World - buffered")
	buf.Flush()

	var str string = "Hello World - string"
	fmt.Println(str[1:])
}
