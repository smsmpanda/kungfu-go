package osrel

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input1 string
var err error

func Gorun12451212312() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err = inputReader.ReadString('n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
