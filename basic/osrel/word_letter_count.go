package osrel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines = 0, 0, 0

func Gorunfsd23() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input , typp S to stop:")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}

		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}

		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len((input)) - 2
	nrwords += len(strings.Fields(input))
	nrlines++
}
