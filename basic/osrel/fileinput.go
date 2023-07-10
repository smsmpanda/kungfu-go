package osrel

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Gorunos123() {
	println("Gorun")
	inputFile, inputError := os.Open("resources/seckill.lua")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}

}
